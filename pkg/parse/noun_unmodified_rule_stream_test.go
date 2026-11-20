//ff:func feature=parse type=rule control=sequence
//ff:what TestNounUnmodifiedRuleStream — read as a stream with the rule "수 and 줄 are the noun where no modifying form stands before them" (and the rule of a modifying form before a dependent noun): 수 after the noun 인구 is the noun (the dependent noun comes first with no rule), and 수 after 알다 with the modifying ㄹ stays the dependent noun
package parse

import "testing"

func TestNounUnmodifiedRuleStream(t *testing.T) {
	boundCases(t, []boundCase{
		{"인구 수는 많다", "수는", "수[의존명사:수]+는[보조사:는]", "수[열쇠:체언]+는[보조사:는]"},
		{"그는 알 수 있다", "수", "수[의존명사:수]", "수[의존명사:수]"},
	}, boundBeforeRule(boundNext, boundNominal), nounUnmodifiedRule(adnBound))
}
