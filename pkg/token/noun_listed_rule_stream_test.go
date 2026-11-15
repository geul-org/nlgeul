//ff:func feature=token type=rule control=sequence
//ff:what TestNounListedRuleStream — read as a stream with the rule "the letters of a dependent noun right before a middle dot are a noun": 시 before ㆍ is the noun (the dependent noun comes first with no rule), and 시 with no middle dot after it stays the dependent noun
package token

import "testing"

func TestNounListedRuleStream(t *testing.T) {
	boundCases(t, []boundCase{
		{"그는 시ㆍ군에 둔다", "시", "시[의존명사:시]", "시[열쇠:체언]"},
		{"그는 시 군에 둔다", "시", "시[의존명사:시]", "시[의존명사:시]"},
	}, nounListedRule(boundNext))
}
