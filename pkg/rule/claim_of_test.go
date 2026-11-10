//ff:func feature=rule type=helper control=sequence
//ff:what TestClaimOf — the name and conclusion a Claim criterion wrote into a result's name; nothing from a name without one
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestClaimOf(t *testing.T) {
	g := toulmin.NewGraph("x")
	g.Rule(ChainNoCase).With(&Claim{Name: "topicSubject", Concludes: "주어"})
	results, err := g.Evaluate(ctxFor(&Word{Chain: Chain{Size: 2}}, Tail{}))
	if err != nil || len(results) != 1 {
		t.Fatal(results, err)
	}
	if name, label := ClaimOf(results[0].Name); name != "topicSubject" || label != "주어" {
		t.Fatal(results[0].Name)
	}
	if name, label := ClaimOf("ChainNoCase"); name != "" || label != "" {
		t.Fatal(name, label)
	}
}
