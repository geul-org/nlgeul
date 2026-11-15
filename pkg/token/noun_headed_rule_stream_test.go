//ff:func feature=token type=rule control=sequence
//ff:what TestNounHeadedRuleStream — read as a stream with the rule "the letters of a dependent noun at the head of a sentence are a noun": 시는 at the head and right after an opening quotation mark is the noun (the dependent noun comes first with no rule), and 시는 after a noun stays the dependent noun
package token

import "testing"

func TestNounHeadedRuleStream(t *testing.T) {
	boundCases(t, []boundCase{
		{"시는 좋다", "시는", "시[의존명사:시]+는[보조사:는]", "시[열쇠:체언]+는[보조사:는]"},
		{"그는 “시는 좋다", "시는", "시[의존명사:시]+는[보조사:는]", "시[열쇠:체언]+는[보조사:는]"},
		{"인구 시는 좋다", "시는", "시[의존명사:시]+는[보조사:는]", "시[의존명사:시]+는[보조사:는]"},
	}, nounHeadedRule(boundNext))
}
