//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestEuDropped — a dropped ㅡ put back where the vowel follows the syllable before it (따라, 치러, 써, 바빠) lives; it dies for a stem the table gives as 르 or 러 irregular (모라, 이러), for the wrong vowel (따러) and for 아 after a stem of one syllable (싸); a tail with no restored ㅡ, or another restored vowel, is never touched; end to end the live ones are read and the dead ones are not
package token

import (
	"strings"
	"testing"
)

func TestEuDropped(t *testing.T) {
	ws := wordsScan("따라", layerCost)
	if ws.euDropped(0, tail{start: 2, base: '르'}) {
		t.Fatal("따라 = 따르+아: 아 after ㅏ")
	}
	if ws.euDropped(0, tail{start: 2}) || ws.euDropped(0, tail{start: 2, base: '라'}) {
		t.Fatal("a tail with no restored ㅡ")
	}
	if !wordsScan("따러", layerCost).euDropped(0, tail{start: 2, base: '르'}) {
		t.Fatal("따러 as 따르+어: 어 after ㅏ")
	}
	if wordsScan("치러", layerCost).euDropped(0, tail{start: 2, base: '르'}) || wordsScan("바빠", layerCost).euDropped(0, tail{start: 2, base: '쁘'}) {
		t.Fatal("치르+어, 바쁘+아")
	}
	if !wordsScan("모라", layerCost).euDropped(0, tail{start: 2, base: '르'}) {
		t.Fatal("모르 is 르 irregular in the table: no 모르+아")
	}
	reo := wordsScan("이러", layerCost)
	reo.ix.lex["이르"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}, irr: irrReo}
	if !reo.euDropped(0, tail{start: 2, base: '르'}) {
		t.Fatal("a stem the table gives as 러 irregular")
	}
	if !wordsScan("싸", layerCost).euDropped(0, tail{start: 1, base: '쓰'}) || wordsScan("써", layerCost).euDropped(0, tail{start: 1, base: '쓰'}) {
		t.Fatal("a stem of one syllable takes 어 only")
	}
	if wordsScan("곰써", layerCost).euDropped(1, tail{start: 2, base: '쓰'}) || !wordsScan("곰싸", layerCost).euDropped(1, tail{start: 2, base: '쓰'}) {
		t.Fatal("a stem of one syllable after a boundary")
	}
	if wordsScan("모로", layerCost).euDropped(0, tail{start: 2, base: '르'}) {
		t.Fatal("a syllable not written with ㅏ or ㅓ")
	}

	r, _ := New(Options{Words: "testdata/words.tsv", Asleep: true})
	r.ix.lex["따르"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["치르"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["쓰"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["바쁘"] = word{kinds: kindAdj, count: 50, by: [6]int{0, 0, 50}}
	for in, want := range map[string]string{"따라": "따르[열쇠:동사]+(아)", "치러": "치르[열쇠:동사]+(어)", "써": "쓰[열쇠:동사]+(어)", "바빠": "바쁘[열쇠:형용사]+(아)"} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	for in, dead := range map[string]string{"모라": "모르[", "따러": "따르[", "싸": "쓰["} {
		if got := strings.Join(r.Lines(in, false), "\n"); strings.Contains(got, dead) {
			t.Errorf("%s: a dropped ㅡ put back where the stem does not drop it\n%s", in, got)
		}
	}
}
