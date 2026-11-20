//ff:func feature=parse type=rule control=sequence
//ff:what TestNounUnheldRuleStream — read as a stream with the rule "the letters of a dependent noun with no modifier and no nominal before them are a noun": 시 after the adverb 또는 is the noun (the dependent noun comes first with no rule), and 시 after the noun 인구 stays the dependent noun
package parse

import "testing"

func TestNounUnheldRuleStream(t *testing.T) {
	boundCases(t, []boundCase{
		{"보고 또는 시는 많다", "시는", "시[의존명사:시]+는[보조사:는]", "시[열쇠:체언]+는[보조사:는]"},
		{"보고 인구 시는 많다", "시는", "시[의존명사:시]+는[보조사:는]", "시[의존명사:시]+는[보조사:는]"},
	}, nounUnheldRule(unheldBound))
}
