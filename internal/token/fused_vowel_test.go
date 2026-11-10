//ff:func feature=token type=reader control=sequence
//ff:what TestFusedVowel — 봐 and 해 always; 가 only with more of the tail; never with a final consonant
package token

import "testing"

func TestFusedVowel(t *testing.T) {
	if !fusedVowel('봐', false) || !fusedVowel('해', false) || fusedVowel('가', false) || !fusedVowel('가', true) || fusedVowel('간', true) || fusedVowel('기', true) {
		t.Fatal("fusedVowel")
	}
}
