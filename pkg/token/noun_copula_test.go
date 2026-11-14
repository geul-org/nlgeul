//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNounCopula — a closing right after 테 (갈테니), fused under it (올텐데) or one that takes only the copula's stem (할테야) is read after 터 with the stem (이) in it: the stem counted as the copula and leaving a trace in the letter, the tail ending in the closing, the letter as the only stem, 터 shown, the graph's name as the label, and the graph and facts kept to be asked again on each candidate; a unit read as attaching takes the modifying form to stand before it; no candidate where no modifying closing ends before 테 (also inside a unit read as attaching), before a closing the copula does not take, before ㅆ fused into its vowel, after another letter (and 텐 with no consonant read under it), for a tail with no closing, a closing that holds the copula, a tail that asks for a changed stem, or nothing before it, and only the graph's misses are judged and traced with their reasons
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounCopula(t *testing.T) {
	adn := func(text string, at int) *scan {
		sc := scanOf(text)
		sc.adnEnds[at] = true
		return sc
	}
	lone := func(text string) *scan {
		sc := scanOf(text)
		sc.attach = true
		return sc
	}
	ni := tail{start: 2, pieces: []piece{{text: "니", names: "맺음:니"}}, close: &entry{name: "맺음:니"}, noYo: true}
	sc := adn("갈테니", 1)
	stem := piece{"(이)", copulaName, 1, sc.ix.pref(copulaCount, "이"), rule.TailCopula, false, ""}
	got := sc.nounCopula(ni)
	if len(got) != 1 || got[0].start != 2 || len(got[0].pieces) != 2 || got[0].pieces[0] != stem || got[0].pieces[1] != ni.pieces[0] || got[0].close != ni.close || !got[0].noYo || len(ni.pieces) != 1 {
		t.Fatal(got)
	}
	if g := got[0]; g.only != "테" || g.shown != "터" || g.label != "줄임:테" || g.again != sc.ix.nouns[0] || g.facts == nil || g.facts.Form != "테" || !g.facts.HasStem || !g.facts.Adn || g.facts.After != "맺음:니" {
		t.Fatal("the noun's letter, the noun, its graph and its facts", g)
	}
	if sc.again(got[0], false) || !sc.again(got[0], true) {
		t.Fatal("asked again on a candidate, it stands only where the candidate's last word ends in a modifying closing")
	}
	sc = adn("올텐데", 1)
	got = sc.nounCopula(tail{start: 2, base: '테', pieces: []piece{{text: "(ㄴ)데", names: "맺음:ㄴ데"}}, close: &entry{name: "맺음:ㄴ데"}})
	if len(got) != 1 || got[0].start != 2 || got[0].only != "텐" || got[0].shown != "터" || got[0].base != 0 || got[0].pieces[0] != stem || got[0].pieces[1].text != "(ㄴ)데" || got[0].facts.After != "맺음:ㄴ데" {
		t.Fatal("ㄴ데 fused under the noun's letter", got)
	}
	sc = adn("할테야", 1)
	got = sc.nounCopula(tail{start: 2, pieces: []piece{{text: "야", names: "맺음:야"}}, close: &entry{name: "맺음:야", needs: true}, only: "아니"})
	if len(got) != 1 || got[0].only != "테" || got[0].pieces[0] != stem || got[0].pieces[1].names != "맺음:야" {
		t.Fatal("야, which takes only the copula's stem (or 아니), after the stem in 테", got)
	}
	head := lone("테니까")
	if got = head.nounCopula(tail{start: 1, pieces: []piece{{text: "니까", names: "맺음:니까"}}, close: &entry{name: "맺음:니까"}}); len(got) != 1 || got[0].only != "테" || !got[0].facts.Adn {
		t.Fatal("a unit read as attaching takes the modifying form to stand in the unit before", got)
	}
	nde := []piece{{text: "(ㄴ)데", names: "맺음:ㄴ데"}}
	for _, c := range []struct {
		sc  *scan
		t   tail
		why string
	}{
		{scanOf("나테니"), ni, whyNoAdn},
		{lone("곰테니"), ni, whyNoAdn},
		{adn("할테자", 1), tail{start: 2, pieces: []piece{{text: "자", names: "맺음:자"}}, close: &entry{name: "맺음:자"}}, whyBarred},
		{adn("할텠다", 1), tail{start: 2, base: '테', pieces: []piece{{text: "(었)", names: "선어말:ㅆ"}, {text: "다", names: "맺음:다"}}, close: &entry{name: "맺음:다"}}, whyFusedIn},
		{adn("갈가니", 1), ni, ""},
		{adn("갈갠데", 1), tail{start: 2, base: '개', pieces: nde, close: &entry{name: "맺음:ㄴ데"}}, ""},
		{adn("갈텐데", 1), tail{start: 2, pieces: []piece{{text: "데", names: "맺음:데"}}, close: &entry{name: "맺음:데"}}, ""},
		{adn("갈테니", 1), tail{start: 2, pieces: ni.pieces}, ""},
		{adn("할테예요", 1), tail{start: 2, pieces: []piece{{text: "예요", names: "맺음:예요"}}, close: &entry{name: "맺음:예요", holds: true}}, ""},
		{adn("갈테니", 1), tail{start: 2, pieces: ni.pieces, close: ni.close, irr: irrS}, ""},
		{lone("니까"), tail{start: 0, pieces: []piece{{text: "니까", names: "맺음:니까"}}, close: &entry{name: "맺음:니까"}}, ""},
	} {
		got := c.sc.nounCopula(c.t)
		notes := strings.Join(*c.sc.trace, "\n")
		if len(got) != 0 || (c.why == "" && notes != "") || !strings.Contains(notes, c.why) {
			t.Errorf("%s at %d: %v\n%s", string(c.sc.s), c.t.start, got, notes)
		}
	}
}
