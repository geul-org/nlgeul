//ff:func feature=token type=rule control=sequence
//ff:what TestNounHeadedRuleStream — read as a stream with the rule "the letters of a dependent noun at the head of a sentence or right before a middle dot are a noun": 시는 at the head and right after an opening quotation mark, and 시 before ㆍ, is the noun (the dependent noun comes first with no rule), and 시는 after a noun and 시 with no middle dot after it stay the dependent noun
package token

import "testing"

func TestNounHeadedRuleStream(t *testing.T) {
	boundCases(t, []boundCase{
		{"시는 좋다", "시는", "시[의존명사:시]+는[보조사:는]", "시[열쇠:체언]+는[보조사:는]"},
		{"그는 “시는 좋다", "시는", "시[의존명사:시]+는[보조사:는]", "시[열쇠:체언]+는[보조사:는]"},
		{"인구 시는 좋다", "시는", "시[의존명사:시]+는[보조사:는]", "시[의존명사:시]+는[보조사:는]"},
		{"그는 시ㆍ군에 둔다", "시", "시[의존명사:시]", "시[열쇠:체언]"},
		{"그는 시 군에 둔다", "시", "시[의존명사:시]", "시[의존명사:시]"},
	}, nounHeadedRule(boundNext))
}
