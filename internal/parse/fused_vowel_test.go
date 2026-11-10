//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestFusedVowel — 봐 줘 져 돼 해 always; 가 서 내 세 only with more of the tail; never with a final consonant or another vowel
package parse

import "testing"

func TestFusedVowel(t *testing.T) {
	for _, r := range "봐줘져돼해" {
		if !fusedVowel(r, false) {
			t.Errorf("%c", r)
		}
	}
	for _, r := range "가서내세" {
		if fusedVowel(r, false) || !fusedVowel(r, true) {
			t.Errorf("%c", r)
		}
	}
	for _, r := range "봤기우A" {
		if fusedVowel(r, true) {
			t.Errorf("%c", r)
		}
	}
}
