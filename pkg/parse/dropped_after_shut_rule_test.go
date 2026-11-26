//ff:func feature=parse type=helper control=sequence
//ff:what TestDroppedAfterShutRule — the rule is named as the trace prints it, adds bonusShortTaken, and speaks of a candidate with the dropped copula stem
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDroppedAfterShutRule(t *testing.T) {
	r := droppedAfterShutRule()
	if r.name != ruleDropShut || r.bonus != bonusShortTaken || r.graph == nil {
		t.Fatalf("%+v", r)
	}
	if !r.about(rule.Near{Stem: "", Last: "맺음:다", Marks: "서술격:이(빠짐)+맺음:다"}) || r.about(rule.Near{Stem: "", Last: "맺음:다", Marks: "맺음:다"}) {
		t.Fatal("about")
	}
}
