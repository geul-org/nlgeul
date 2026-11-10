//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestFusedVowel — 봐 줘 져 돼 해 always; 가 서 내 세 only with more of the tail; never with a final consonant or another vowel
package rule

import "testing"

func TestFusedVowel(t *testing.T) {
	for _, r := range "봐줘져돼해" {
		if !FusedVowel(r, false) {
			t.Errorf("%c", r)
		}
	}
	for _, r := range "가서내세" {
		if FusedVowel(r, false) || !FusedVowel(r, true) {
			t.Errorf("%c", r)
		}
	}
	for _, r := range "봤기우A" {
		if FusedVowel(r, true) {
			t.Errorf("%c", r)
		}
	}
}
