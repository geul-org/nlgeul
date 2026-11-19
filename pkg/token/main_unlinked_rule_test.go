//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestMainUnlinkedRule — the rule on its own, a candidate judged with the words around it given: it speaks of a candidate whose stem is an auxiliary given read as its main predicate (있 the adjective, 보 the verb; not 있 the verb, not the auxiliary reading, not a stem not given), its name and its bonus are bonusMainFirst, it stands when no word or a word that ends in another closing is before, and it falls with the reasons 앞 어절이 연결 꼴(어/아, 고)로 끝난다 after a linking form and 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝난다 after a question form
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMainUnlinkedRule(t *testing.T) {
	brs, _ := auxLinkedLast(t, "있다", false, false)
	adj := auxLinkedPick(t, brs, "있[열쇠:형용사]+다[맺음:다]")
	r := mainUnlinkedRule([]string{"있", "보"})
	if r.name != ruleMainUnlinked || r.bonus != bonusMainFirst || r.waits != nil || r.graph == nil {
		t.Fatalf("%+v", r)
	}
	for near, want := range map[rule.Near]bool{
		{Stem: "있", Kind: "열쇠:형용사"}:  true,
		{Stem: "보", Kind: "열쇠:동사"}:   true,
		{Stem: "있", Kind: "열쇠:동사"}:   false,
		{Stem: "보", Kind: "열쇠:형용사"}:  false,
		{Stem: "있", Kind: "보조 용언:있"}: false,
		{Stem: "주", Kind: "열쇠:동사"}:   false,
		{Stem: "집", Kind: "열쇠:체언"}:   false,
	} {
		if r.about(near) != want {
			t.Fatalf("%+v: about %v", near, !want)
		}
	}
	st := sentenceStream(t, "", true, r)
	line := func(v, why string) string {
		return "  · 있[열쇠:형용사]+다[맺음:다]\t" + ruleMainUnlinked + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		before []rule.Near
		add    float64
		note   string
	}{
		{nil, bonusMainFirst, line("1.00", "통과")},
		{[]rule.Near{{Last: "목적격:를"}}, bonusMainFirst, line("1.00", "통과")},
		{[]rule.Near{{Last: "맺음:어"}}, 0, line("-1.00", "기각: 앞 어절이 연결 꼴(어/아, 고)로 끝난다")},
		{[]rule.Near{{Last: "맺음:어(합침)"}}, 0, line("-1.00", "기각: 앞 어절이 연결 꼴(어/아, 고)로 끝난다")},
		{[]rule.Near{{Last: "맺음:고"}}, 0, line("-1.00", "기각: 앞 어절이 연결 꼴(어/아, 고)로 끝난다")},
		{[]rule.Near{{Last: "맺음:나"}}, 0, line("-1.00", "기각: 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝난다")},
	} {
		add, _, notes := st.judgeCandidate(adj, rule.Around{Self: nearOf(adj), Before: c.before})
		if add != c.add || len(notes) != 1 || notes[0] != c.note {
			t.Fatalf("%v: %v %q", c.before, add, notes)
		}
	}
}
