//ff:func feature=token type=rule control=sequence
//ff:what TestBoundBeforeRuleStream — read as a stream with the rule "a modifying form stands right before a dependent noun": 알 before 수 is 알다 with the modifying ㄹ (the noun 알 comes first with no rule), and 알 before 사과 stays the noun
package token

import "testing"

func TestBoundBeforeRuleStream(t *testing.T) {
	boundCases(t, []boundCase{
		{"그는 알 수 있다", "알", "알[열쇠:체언]", "알[열쇠:동사]+(ㄹ)[꾸밈:ㄹ]"},
		{"그는 알 인구", "알", "알[열쇠:체언]", "알[열쇠:체언]"},
	}, boundBeforeRule(boundNext, boundNominal))
}
