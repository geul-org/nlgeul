//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestMalLinkedRule — the rule on its own, a candidate judged with the words around it given: it speaks only of the stem 말 read as the verb (not a noun 말, not 마라 the noun), its name and its bonus are bonusShortTaken and it does not wait, it stands (통과) when the word before ends in 지 or 고, and it falls with the reason 앞 어절이 지·고로 끝나지 않는다 after another closing or no word
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMalLinkedRule(t *testing.T) {
	words, _ := linkBeforeWords(t, "집 마라", false)
	mal := auxLinkedPick(t, words[1], "말[열쇠:동사]+(아)라[맺음:어라(합침)]")
	noun := auxLinkedPick(t, words[1], "마라[열쇠:체언]")
	r := malLinkedRule()
	if r.name != ruleMalLinked || r.bonus != bonusShortTaken || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	if !r.about(nearOf(mal)) || r.about(nearOf(noun)) || r.about(rule.Near{Stem: "말", Kind: "열쇠:체언"}) {
		t.Fatal("it speaks of the verb 말 only")
	}
	st := sentenceStream(t, "", true, r)
	line := func(v, why string) string {
		return "  · 말[열쇠:동사]+(아)라[맺음:어라(합침)]\t" + ruleMalLinked + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		before []rule.Near
		add    float64
		note   string
	}{
		{[]rule.Near{{Stem: "울", Last: "맺음:지"}}, bonusShortTaken, line("1.00", "통과")},
		{[]rule.Near{{Stem: "울", Last: "맺음:고"}}, bonusShortTaken, line("1.00", "통과")},
		{[]rule.Near{{Stem: "울", Last: "맺음:어"}}, 0, line("-1.00", "기각: 앞 어절이 지·고로 끝나지 않는다")},
		{[]rule.Near{{Stem: "집", Kind: "열쇠:체언"}}, 0, line("-1.00", "기각: 앞 어절이 지·고로 끝나지 않는다")},
		{nil, 0, line("-1.00", "기각: 앞 어절이 지·고로 끝나지 않는다")},
	} {
		add, _, notes := st.judgeCandidate(mal, rule.Around{Self: nearOf(mal), Before: c.before})
		if add != c.add || len(notes) != 1 || notes[0] != c.note {
			t.Fatalf("%v: %v %q", c.before, add, notes)
		}
	}
}
