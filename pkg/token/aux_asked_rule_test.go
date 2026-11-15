//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAuxAskedRule — the rule on its own, a candidate judged with the words around it given: it speaks only of the candidate read as the auxiliary 보 (보조 용언:보, not 보조 용언:있, not the main verb 보), its name and its bonus are bonusSentence, it stands after -나, -ㄴ가/-는가/-은가 and -ㄹ까/-을까, and it falls with the reason 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝나지 않는다 after a linking form, another closing or no word
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAuxAskedRule(t *testing.T) {
	brs, _ := auxLinkedLast(t, "보다", false, false)
	aux := auxLinkedPick(t, brs, "보[보조 용언:보]+다[맺음:다]")
	main := auxLinkedPick(t, brs, "보[열쇠:동사]+다[맺음:다]")
	r := auxAskedRule()
	if r.name != ruleAuxAsked || r.bonus != bonusSentence || r.waits != nil || r.graph == nil {
		t.Fatalf("%+v", r)
	}
	if !r.about(nearOf(aux)) || r.about(nearOf(main)) || r.about(rule.Near{Kind: "보조 용언:있"}) {
		t.Fatal("it speaks of the auxiliary 보 only")
	}
	st := sentenceStream(t, "", true, r)
	line := func(v, why string) string {
		return "  · 보[보조 용언:보]+다[맺음:다]\t" + ruleAuxAsked + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		before []rule.Near
		add    float64
		note   string
	}{
		{[]rule.Near{{Last: "맺음:나"}}, bonusSentence, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:ㄴ가"}}, bonusSentence, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:는가"}}, bonusSentence, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:은가"}}, bonusSentence, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:ㄹ까"}}, bonusSentence, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:을까"}}, bonusSentence, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:어"}}, 0, line("-1.00", "기각: 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝나지 않는다")},
		{[]rule.Near{{Last: "맺음:면"}}, 0, line("-1.00", "기각: 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝나지 않는다")},
		{nil, 0, line("-1.00", "기각: 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝나지 않는다")},
	} {
		add, notes := st.judgeCandidate(aux, rule.Around{Self: nearOf(aux), Before: c.before})
		if add != c.add || len(notes) != 1 || notes[0] != c.note {
			t.Fatalf("%v: %v %q", c.before, add, notes)
		}
	}
}
