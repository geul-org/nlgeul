//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestSentenceRebuttal — its claim stands for the candidates the list names and is the rebuttal's (the result's name carries the rule's name and 반박): evaluated with a candidate in the list the verdict is above 0, with one not in the list it is not; an exception put on it defeats it
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSentenceRebuttal(t *testing.T) {
	g, claim := sentenceRebuttal("문장:시험 반박", rule.Listed, listed("줄기", rule.ListSelfStem, []string{"있"}))
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
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyAround, rule.Around{Self: rule.Near{Stem: "있"}})
	results, err := g.Evaluate(ctx)
	if err != nil || len(results) != 1 {
		t.Fatal(results, err)
	}
	if name, label := rule.ClaimOf(results[0].Name); name != "문장:시험 반박" || label != rebutLabel {
		t.Fatal(name, label)
	}
	except(g, rule.Listed, "뒤가 되다", listed("되", rule.ListNextStem, []string{"되"})).Attacks(claim)
	ctx = toulmin.NewContext()
	ctx.Set(rule.KeyAround, rule.Around{Self: rule.Near{Stem: "있"}, After: []rule.Near{{Stem: "되"}}})
	if v, why := evaluate(g, ctx, true); v > 0 || len(why) != 1 || why[0] != "뒤가 되다" {
		t.Fatal(v, why)
	}
}
