//ff:func feature=parse type=rule control=sequence
//ff:what TestClaim — a claim node is declared under its name with its criteria first and the Claim that names it; the result of evaluating it carries name and conclusion
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestClaim(t *testing.T) {
	g := toulmin.NewGraph("x")
	c := map[string]*toulmin.Rule{}
	claim(g, c, "objectCase", labelObj, rule.ChainCase, &rule.Bits{Of: "case", Mask: caseObj})
	claim(g, c, "topicObject", labelObj, rule.ChainNoCase)
	results, err := g.Evaluate(ctxFor(readWord("밥을", false, &sentence{}), tailFact{}))
	if err != nil || len(c) != 2 || len(results) != 1 {
		t.Fatal(results, err)
	}
	if name, label := rule.ClaimOf(results[0].Name); name != "objectCase" || label != labelObj {
		t.Fatal(results[0].Name)
	}
}
