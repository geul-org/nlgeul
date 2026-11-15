//ff:func feature=token type=rule control=sequence
//ff:what TestNounModifiedRuleStream — read as a stream with the rule "a noun stands right after a modifier": 보고 after 따르다 with the modifying ㄴ is the noun (보다 with 고 comes first with no rule), and 보고 after an adverb stays 보다 with 고
package token

import "testing"

func TestNounModifiedRuleStream(t *testing.T) {
	boundCases(t, []boundCase{
		{"따른 보고 또는", "보고", "보[보조 용언:보]+고[맺음:고]", "보고[열쇠:체언]"},
		{"또는 보고 따른", "보고", "보[보조 용언:보]+고[맺음:고]", "보[보조 용언:보]+고[맺음:고]"},
	}, nounModifiedRule(modifiedSkip))
}
