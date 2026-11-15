//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestDeterminerOneRule — the rule on its own, a candidate judged with the words around it given: it speaks only of 한 read as the determiner of the closed set and written apart (not glued, not 열쇠:관형사, not 하다 with ㄴ), its name and its bonus are bonusOne, it waits for nothing; it stands at the head of a sentence and after a bare noun, 의 or a modifying ㄴ, and it falls after each form of haTakes (an object, 로/으로, -게, -도록, -어야, -려(고), -고자, quoting) with the reason whyHaTakes and after the modifying 는 with the reason whyNeun
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDeterminerOneRule(t *testing.T) {
	brs, _ := oneCandidates(t, "한", false, false)
	det := auxLinkedPick(t, brs, "한[관형사:한]")
	ha := auxLinkedPick(t, brs, "하[열쇠:동사]+(ㄴ)[꾸밈:ㄴ]")
	r := determinerOneRule()
	if r.name != ruleDeterminerOne || r.bonus != bonusOne || r.waits != nil || r.graph == nil {
		t.Fatalf("%+v", r)
	}
	glued := nearOf(det)
	glued.Glued = true
	if !r.about(nearOf(det)) || r.about(glued) || r.about(nearOf(ha)) || r.about(rule.Near{Stem: "한", Kind: "열쇠:관형사"}) {
		t.Fatal("it speaks of the determiner 한 written apart only")
	}
	st := sentenceStream(t, "", true, r)
	line := func(v, why string) string {
		return "  · 한[관형사:한]\t" + ruleDeterminerOne + "\t" + v + "\t" + why
	}
	stands, takes, neun := line("1.00", "통과"), line("-1.00", "기각: "+whyHaTakes), line("-1.00", "기각: "+whyNeun)
	prev := func(last string) []rule.Near { return []rule.Near{{Last: last}} }
	cases := []struct {
		before []rule.Near
		add    float64
		note   string
	}{
		{nil, bonusOne, stands},
		{[]rule.Near{{Stem: "사람", Kind: "열쇠:체언"}}, bonusOne, stands},
		{prev("관형격:의"), bonusOne, stands},
		{prev("꾸밈:ㄴ"), bonusOne, stands},
		{prev("꾸밈:는"), 0, neun},
		{prev("목적격:을"), 0, takes}, {prev("목적격:를"), 0, takes}, {prev("부사격:로"), 0, takes}, {prev("부사격:으로"), 0, takes},
		{prev("맺음:게"), 0, takes}, {prev("맺음:도록"), 0, takes}, {prev("맺음:어야"), 0, takes}, {prev("맺음:여야"), 0, takes},
		{prev("맺음:려"), 0, takes}, {prev("맺음:으려"), 0, takes}, {prev("맺음:려고"), 0, takes}, {prev("맺음:으려고"), 0, takes},
		{prev("맺음:고자"), 0, takes}, {prev("맺음:다고"), 0, takes}, {prev("맺음:ㄴ다고"), 0, takes}, {prev("맺음:는다고"), 0, takes},
		{prev("맺음:라고"), 0, takes}, {prev("맺음:냐고"), 0, takes}, {prev("맺음:자고"), 0, takes},
	}
	for _, c := range cases {
		add, notes := st.judgeCandidate(det, rule.Around{Self: nearOf(det), Before: c.before})
		if add != c.add || len(notes) != 1 || notes[0] != c.note {
			t.Fatalf("%v: %v %q", c.before, add, notes)
		}
	}
}
