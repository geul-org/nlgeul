//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddStems — two graphs, the new word's of three edges without a word table and of six with one, the restored syllable's of nine; a new word's verdict is 1/3; a new word whose letters read as markers does not stand whole, and without a word table stands with markers after it; with a word table a stem in ㅆ before markers and a word the tables know only as kinds that take no predicate tail before one do not stand, while a predicate's stem too and a tail of particles do; the restored syllable does not stand right after a letter that is no Hangul (but for 하, 되, 이 and at the stem's head), for a stem kept as 르 or 러 irregular, for 어/아 that does not fit the syllable before, nor for 아 under a stem of one syllable; 하 as it stands in the letters takes no 어/아 that leaves no trace, but one that left its trace (해) and any other syllable (가, 서) stand, with the reason on its own node; each is turned down for the reason on its own node
package token

import (
	"math"
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddStems(t *testing.T) {
	bare, lexed := &index{}, &index{lexed: true}
	bare.addStems()
	lexed.addStems()
	if bare.graphs != 2 || bare.edges != 3+9 || lexed.graphs != 2 || lexed.edges != 6+9 || bare.stem == nil || bare.restore == nil || lexed.stem == nil || lexed.restore == nil {
		t.Fatal(bare.graphs, bare.edges, lexed.graphs, lexed.edges)
	}
	if math.Abs(bare.unknown-1.0/3) > 1e-9 || math.Abs(lexed.unknown-1.0/3) > 1e-9 {
		t.Fatal("a new word's verdict", bare.unknown, lexed.unknown)
	}
	for _, c := range []struct {
		g      *index
		p      rule.Piece
		stands bool
		why    string
	}{
		{bare, rule.Piece{Form: "에", Tail: rule.TailBare, Marked: true}, false, whyMarked},
		{bare, rule.Piece{Form: "에", Tail: rule.TailCase, Marked: true}, true, whyUntold},
		{bare, rule.Piece{Form: "곰졌", Tail: rule.TailPred}, true, ""},
		{bare, rule.Piece{Form: "사과", Tail: rule.TailPred, Stem: kindNominal}, true, ""},
		{lexed, rule.Piece{Form: "곰", Tail: rule.TailPred}, true, ""},
		{lexed, rule.Piece{Form: "에", Tail: rule.TailBare, Marked: true}, false, whyMarked},
		{lexed, rule.Piece{Form: "에", Tail: rule.TailCase, Marked: true}, false, whyMarked},
		{lexed, rule.Piece{Form: "곰졌", Tail: rule.TailPred}, false, whyCoda},
		{lexed, rule.Piece{Form: "곰졌", Tail: rule.TailAux}, false, whyCoda},
		{lexed, rule.Piece{Form: "곰졌", Tail: rule.TailBare}, true, ""},
		{lexed, rule.Piece{Form: "곰갓", Tail: rule.TailPred}, true, ""},
		{lexed, rule.Piece{Form: "사과", Tail: rule.TailPred, Stem: kindNominal}, false, whyKnown},
		{lexed, rule.Piece{Form: "어떤", Tail: rule.TailAdj, Stem: kindDeterminer}, false, whyKnown},
		{lexed, rule.Piece{Form: "물", Tail: rule.TailPred, Stem: kindNominal | kindVerb}, true, whyPredicate},
		{lexed, rule.Piece{Form: "사과", Tail: rule.TailCase, Stem: kindNominal}, true, whyNoPredTail},
	} {
		v, why := verdictOf(c.g.stem, c.p, "", true)
		if (v > 0) != c.stands || (c.why != "" && !slices.Contains(why, c.why)) || (c.why == "" && len(why) != 0) || (c.stands && math.Abs(v-1.0/3) > 1e-9) {
			t.Errorf("new word %+v (with a word table: %v): verdict %v, reasons %v", c.p, c.g.lexed, v, why)
		}
	}
	for _, c := range []struct {
		p      rule.Piece
		stands bool
		why    string
	}{
		{rule.Piece{Base: "고", HasStem: true, Prev: '사'}, true, ""},
		{rule.Piece{Base: "고", HasStem: true, Prev: 'i'}, false, whyNoHangul},
		{rule.Piece{Base: "고", HasStem: true, Prev: '2'}, false, whyNoHangul},
		{rule.Piece{Base: "하", HasStem: true, Prev: 'h'}, true, whyMakes},
		{rule.Piece{Base: "되", HasStem: true, Prev: 'e'}, true, whyMakes},
		{rule.Piece{Base: "이", HasStem: true, Prev: 'i'}, true, whyMakes},
		{rule.Piece{Base: "고"}, true, whyStemHead},
		{rule.Piece{Base: "르", HasStem: true, Prev: '따', Form: "아"}, true, whyLonger},
		{rule.Piece{Base: "르", HasStem: true, Prev: '모', Form: "아", Keeps: true}, false, whyKeeps},
		{rule.Piece{Base: "르", HasStem: true, Prev: '따', Form: "어"}, false, whyClash},
		{rule.Piece{Base: "르", HasStem: true, Prev: '치', Form: "어"}, true, ""},
		{rule.Piece{Base: "르", HasStem: true, Prev: '치', Form: "아"}, false, whyClash},
		{rule.Piece{Base: "쓰", Form: "아"}, false, whyOne},
		{rule.Piece{Base: "쓰", Form: "어"}, true, whyStemHead},
		{rule.Piece{Base: "쁘", HasStem: true, Prev: '바', Form: "아"}, true, whyLonger},
	} {
		v, why := verdictOf(bare.restore, c.p, "", true)
		if (v > 0) != c.stands || (c.why != "" && !slices.Contains(why, c.why)) || (c.why == "" && len(why) != 0) || verdict(lexed.restore, c.p) != v {
			t.Errorf("restored syllable %+v: verdict %v, reasons %v", c.p, v, why)
		}
	}
}
