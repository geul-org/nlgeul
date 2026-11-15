//ff:func feature=token type=rule control=sequence
//ff:what TestBoundAfterRuleStream — read as a stream with the rule "a dependent noun stands right after a modifying form": 수도 after 가다 with the modifying ㄹ is the dependent noun 수 with 도 (the noun 수도 comes first with no rule), and 수도 after the noun 인구 stays the noun
package token

import "testing"

func TestBoundAfterRuleStream(t *testing.T) {
	boundCases(t, []boundCase{
		{"갈 수도 있다", "수도", "수도[열쇠:체언]", "수[의존명사:수]+도[보조사:도]"},
		{"인구 수도 많다", "수도", "수도[열쇠:체언]", "수도[열쇠:체언]"},
	}, boundAfterRule(boundNominal))
}
