//ff:func feature=token type=rule control=sequence
//ff:what TestUnmodifiedRule — the rule is named as the trace prints it, adds bonusShortTaken (what the rebuttal "one syllable long" alone takes from lambda, about 3.33), and speaks of a noun of the table that ends in 인 or 일 (adnNoun), for which it waits for the word after it, and of nothing else
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestUnmodifiedRule(t *testing.T) {
	r := unmodifiedRule()
	if r.name != ruleUnmodified || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	if math.Abs(r.bonus-Lambda*math.Log10(1/0.6)) > 1e-9 || r.bonus != bonusShortTaken {
		t.Fatal("bonus", r.bonus)
	}
	noun := rule.Near{Stem: "감정인", Kind: "열쇠:체언"}
	if !r.about(noun) || !r.waits(noun) {
		t.Fatal("감정인")
	}
	copula := rule.Near{Stem: "감정", Kind: "열쇠:체언", Last: "꾸밈:ㄴ", Marks: "서술격:이+꾸밈:ㄴ", Pred: true}
	if r.about(copula) || r.waits(copula) {
		t.Fatal("감정+이+ㄴ")
	}
}
