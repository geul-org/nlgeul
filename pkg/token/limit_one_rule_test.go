//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestLimitOneRule — the rule on its own, a candidate judged with the words around it given: it speaks only of 한 read as a nominal of the word table (not the determiner 한, not 하 read as a nominal), its name and its bonus are bonusOverCount, it waits for nothing; it stands after the modifying 는 and falls with the reason whyNoNeun after the modifying ㄴ, an object or no word
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLimitOneRule(t *testing.T) {
	brs, _ := oneCandidates(t, "한", false, false)
	noun := auxLinkedPick(t, brs, "한[열쇠:체언]")
	det := auxLinkedPick(t, brs, "한[관형사:한]")
	r := limitOneRule()
	if r.name != ruleLimitOne || r.bonus != bonusOverCount || r.waits != nil || r.graph == nil {
		t.Fatalf("%+v", r)
	}
	if !r.about(nearOf(noun)) || r.about(nearOf(det)) || r.about(rule.Near{Stem: "하", Kind: "열쇠:체언"}) {
		t.Fatal("it speaks of 한 read as a nominal only")
	}
	st := sentenceStream(t, "", true, r)
	line := func(v, why string) string { return "  · 한[열쇠:체언]\t" + ruleLimitOne + "\t" + v + "\t" + why }
	falls := line("-1.00", "기각: "+whyNoNeun)
	for _, c := range []struct {
		before []rule.Near
		add    float64
		note   string
	}{
		{[]rule.Near{{Stem: "있", Last: "꾸밈:는", Pred: true}}, bonusOverCount, line("1.00", "통과")},
		{[]rule.Near{{Stem: "가능하", Last: "꾸밈:ㄴ", Pred: true}}, 0, falls},
		{[]rule.Near{{Stem: "일", Last: "목적격:을"}}, 0, falls},
		{nil, 0, falls},
	} {
		add, notes := st.judgeCandidate(noun, rule.Around{Self: nearOf(noun), Before: c.before})
		if add != c.add || len(notes) != 1 || notes[0] != c.note {
			t.Fatalf("%v: %v %q", c.before, add, notes)
		}
	}
}
