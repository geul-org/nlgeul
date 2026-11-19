//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestMalLinkedRuleParticle — 말다 judged after a word whose markers hold 지 or 고 and end in an auxiliary particle: 말+(아)라 after -지도 and after -고야 stands and adds bonusShortTaken; after a noun with 도 (no 지 or 고) it falls with the reason 앞 어절이 지·고로 끝나지 않는다
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMalLinkedRuleParticle(t *testing.T) {
	mal := "말[열쇠:동사]+(아)라[맺음:어라(합침)]"
	st := sentenceStream(t, "", true, malLinkedRule())
	words, _ := linkBeforeWords(t, "마라", false)
	cand := auxLinkedPick(t, words[0], mal)
	for _, c := range []struct {
		before rule.Near
		add    float64
		note   string
	}{
		{rule.Near{Stem: "울", Kind: "열쇠:동사", Last: "보조사:도", Marks: "맺음:지+보조사:도"}, bonusShortTaken, "1.00\t통과"},
		{rule.Near{Stem: "울", Kind: "열쇠:동사", Last: "보조사:야", Marks: "맺음:고+보조사:야"}, bonusShortTaken, "1.00\t통과"},
		{rule.Near{Stem: "집", Kind: "열쇠:체언", Last: "보조사:도", Marks: "보조사:도"}, 0, "-1.00\t기각: " + whyNotMalLinked},
	} {
		add, _, notes := st.judgeCandidate(cand, rule.Around{Self: nearOf(cand), Before: []rule.Near{c.before}})
		if add != c.add || len(notes) != 1 || notes[0] != "  · "+mal+"\t"+ruleMalLinked+"\t"+c.note {
			t.Fatalf("%+v: %v %q", c.before, add, notes)
		}
	}
}
