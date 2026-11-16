//ff:func feature=token type=rule control=sequence
//ff:what TestGluedShortRule — the rule is named as the trace prints it, adds bonusShortTaken (what the rebuttal "one syllable long" alone takes from lambda, about 3.33), waits for nothing, and speaks of a noun of one syllable with its particle written after it (shortNominal) and of nothing else
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGluedShortRule(t *testing.T) {
	r := gluedShortRule()
	if r.name != ruleGluedShort || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	if math.Abs(r.bonus-Lambda*math.Log10(1/0.6)) > 1e-9 || r.bonus != bonusShortTaken {
		t.Fatal("bonus", r.bonus)
	}
	if !r.about(rule.Near{Stem: "자", Kind: "열쇠:체언", Marks: "관형격:의", Text: "자의"}) {
		t.Fatal("자+의")
	}
	if r.about(rule.Near{Stem: "자의", Kind: "열쇠:체언", Text: "자의"}) {
		t.Fatal("자의 alone")
	}
}
