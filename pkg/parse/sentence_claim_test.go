//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestSentenceClaim — its claim stands for the candidates the list names: evaluated with a candidate in the list the verdict is above 0, with one not in the list it is not, and a trace line printed for a rule built on it carries the name given
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSentenceClaim(t *testing.T) {
	g, claim := sentenceClaim("문장:시험", rule.Listed, listed("줄기", rule.ListSelfStem, []string{"있"}))
	if g == nil || claim == nil {
		t.Fatalf("%v %v", g, claim)
	}
	for stem, want := range map[string]bool{"있": true, "집": false} {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, rule.Around{Self: rule.Near{Stem: stem}})
		if v, _ := evaluate(g, ctx, false); (v > 0) != want {
			t.Fatalf("%q: %v", stem, v)
		}
	}
	brs, _ := auxLinkedLast(t, "있다", false, false)
	br := auxLinkedPick(t, brs, "있[열쇠:형용사]+다[맺음:다]")
	st := sentenceStream(t, "", true, sentenceRule{name: "문장:시험", graph: g, bonus: 1, about: func(rule.Near) bool { return true }})
	add, _, notes := st.judgeCandidate(br, rule.Around{Self: nearOf(br)})
	if add != 1 || len(notes) != 1 || notes[0] != "  · 있[열쇠:형용사]+다[맺음:다]\t문장:시험\t1.00\t통과" {
		t.Fatal(add, notes)
	}
}
