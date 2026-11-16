//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAuxLinkedRule — the rule on its own, a candidate judged with the words around it given: it speaks only of the candidates read as the auxiliaries given (보조 용언:있, not 보조 용언:주 when 주 is not given, not the main verb 있), its name and its bonus are bonusOverCount, it stands (the trace says 통과) when the word before ends in 어/아 or its fused form or in 고, and it falls with the reason 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다 when the word before ends in another closing or no word is before
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAuxLinkedRule(t *testing.T) {
	brs, _ := auxLinkedLast(t, "있다", false, false)
	aux := auxLinkedPick(t, brs, "있[보조 용언:있]+다[맺음:다]")
	main := auxLinkedPick(t, brs, "있[열쇠:동사]+다[맺음:다]")
	r := auxLinkedRule([]string{"있", "보"})
	if r.name != ruleAuxLinked || r.bonus != bonusOverCount || r.waits != nil || r.graph == nil {
		t.Fatalf("%+v", r)
	}
	if !r.about(rule.Near{Kind: "보조 용언:있"}) || !r.about(rule.Near{Kind: "보조 용언:보"}) || r.about(rule.Near{Kind: "보조 용언:주"}) || r.about(nearOf(main)) {
		t.Fatal("it speaks of the auxiliaries given only")
	}
	st := sentenceStream(t, "", true, r)
	line := func(v, why string) string {
		return "  · 있[보조 용언:있]+다[맺음:다]\t" + ruleAuxLinked + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		before []rule.Near
		add    float64
		note   string
	}{
		{[]rule.Near{{Last: "맺음:어"}}, bonusOverCount, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:어(합침)"}}, bonusOverCount, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:고"}}, bonusOverCount, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:면"}}, 0, line("-1.00", "기각: 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다")},
		{[]rule.Near{{Last: "목적격:를"}}, 0, line("-1.00", "기각: 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다")},
		{[]rule.Near{{Stem: "집"}}, 0, line("-1.00", "기각: 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다")},
		{nil, 0, line("-1.00", "기각: 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다")},
	} {
		add, notes := st.judgeCandidate(aux, rule.Around{Self: nearOf(aux), Before: c.before})
		if add != c.add || len(notes) != 1 || notes[0] != c.note {
			t.Fatalf("%v: %v %q", c.before, add, notes)
		}
	}
	if add, notes := st.judgeCandidate(main, rule.Around{Self: nearOf(main), Before: []rule.Near{{Last: "맺음:어"}}}); add != 0 || notes != nil {
		t.Fatal("the main verb is not spoken of", add, notes)
	}
}
