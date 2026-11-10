//ff:func feature=parse type=rule control=sequence
//ff:what TestNewGraph — an exception can be laid on a claim after the graph is built: a rule that knows the word 함께 defeats its adverbial reading (and with it the division 함+께), and no other word's
package parse

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestNewGraph(t *testing.T) {
	g, claims := newGraph()
	if len(claims) != len(chainClaims)+len(tailClaims)+len(wordClaims) {
		t.Fatal(len(claims))
	}
	if got, _ := classify(g, "우리가 함께 간다"); got[1] != "함+께:부사어" {
		t.Fatal(got)
	}
	notParticle := func(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
		return rule.Applies(rule.WordOf(ctx).Text == "함께"), nil
	}
	g.Except(notParticle).Attacks(claims["adverbialCase"])
	got, err := classify(g, "우리가 함께 선생님께 간다")
	if err != nil || got[1] != "함께:미정" || got[2] != "선생님+께:부사어" {
		t.Fatal(got, err)
	}
}
