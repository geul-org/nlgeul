//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestSentenceRebuttalSpecs — the claim of a rebuttal takes every criterion given: rule.AroundIs with its bits and a Not stands for a candidate the fact is not true of and not for one it is
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSentenceRebuttalSpecs(t *testing.T) {
	g, _ := sentenceRebuttal("문장:시험 반박", rule.AroundIs, &rule.Bits{Of: "around", Mask: rule.AroundGlued}, not)
	for glued, want := range map[bool]bool{false: true, true: false} {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, rule.Around{Self: rule.Near{Stem: "각", Glued: glued}})
		if v, _ := evaluate(g, ctx, false); (v > 0) != want {
			t.Fatal(glued, v)
		}
	}
}
