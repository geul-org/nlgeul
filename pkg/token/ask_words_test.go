//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestAskWords — the set 의문 대명사 of the closed words holds the five question pronouns 누구, 무엇, 뭐, 어디 and 언제
package token

import (
	"slices"
	"testing"
)

func TestAskWords(t *testing.T) {
	for _, w := range []string{"누구", "무엇", "뭐", "어디", "언제"} {
		if !slices.Contains(askWords, w) {
			t.Errorf("%s is not in askWords %v", w, askWords)
		}
	}
}
