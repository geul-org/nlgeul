//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAddMisfit — with a word table every particle of the table that the final consonant of the word before it chooses has one more entry of its name, read where the other form of its pair is right (은 after a vowel, 는 after a consonant, 로 after a consonant other than ㄹ) and nowhere else: its graph turns it down where its own form fits and after what is no Hangul syllable, and where it stands a rebuttal lowers it to (1-forceMisfit)/(1+forceMisfit); its gate asks for a nominal of the word table; a particle with no condition (에, 도) has none, and with no word table there is none at all
package parse

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddMisfit(t *testing.T) {
	bare, err := newIndex(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, list := range bare.byKey {
		if misfitsIn(list, "") != 0 {
			t.Fatal("no reading against the norm with no word table")
		}
	}
	ix, err := newIndex(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	low := (1 - forceMisfit) / (1 + forceMisfit)
	for _, c := range []struct {
		name       string
		key        rune
		wrong, fit rune
	}{
		{"보조사:은", '은', '교', '당'},
		{"보조사:는", '는', '당', '교'},
		{"목적격:을", '을', '과', '물'},
		{"목적격:를", '를', '물', '과'},
		{"주격:이", '이', '교', '당'},
		{"보격:가", '가', '당', '교'},
		{"부사격:와", '와', '당', '교'},
		{"부사격:로", '로', '당', '물'},
		{"부사격:으로", '로', '물', '당'},
		{"보조사:야", '야', '당', '교'},
		{"보조사:는커녕", '녕', '당', '교'},
	} {
		if misfitsIn(ix.byKey[c.key], c.name) != 1 {
			t.Fatalf("%s: one reading against the norm, got %d", c.name, misfitsIn(ix.byKey[c.key], c.name))
		}
		e := misfitOf(ix.byKey[c.key], c.name)
		if e.gate == nil || e.gate.name != misfitGate || e.class != classParticle {
			t.Fatal(c.name, e.gate)
		}
		if v, _ := verdictOf(e.g, rule.Piece{Form: e.form, HasStem: true, Prev: c.wrong, Next: -1}, "", true); math.Abs(v-low) > 1e-9 {
			t.Errorf("%s after %q: %v, want %v", c.name, c.wrong, v, low)
		}
		if v, _ := verdictOf(e.g, rule.Piece{Form: e.form, HasStem: true, Prev: c.fit, Next: -1}, "", true); v > 0 {
			t.Errorf("%s after %q stands where its own form fits: %v", c.name, c.fit, v)
		}
		if v, _ := verdictOf(e.g, rule.Piece{Form: e.form, HasStem: true, Prev: 'A', Next: -1}, "", true); v > 0 {
			t.Errorf("%s after a Latin letter: %v", c.name, v)
		}
	}
	if misfitsIn(ix.byKey['에'], "부사격:에") != 0 || misfitsIn(ix.byKey['도'], "보조사:도") != 0 {
		t.Fatal("a particle that attaches anywhere has no reading against the norm")
	}
	if v := ix.judged(misfitOf(ix.byKey['은'], "보조사:은").gate.g, rule.Piece{Form: "나무", HasStem: true, Next: -1}, "나무"); v > 0 {
		t.Fatal("a word the table does not hold passes the gate", v)
	}
	if v := ix.judged(misfitOf(ix.byKey['은'], "보조사:은").gate.g, rule.Piece{Form: "학교", HasStem: true, Next: -1, Stem: kindNominal}, "학교"); v <= 0 {
		t.Fatal("a nominal of the table does not pass the gate", v)
	}
}
