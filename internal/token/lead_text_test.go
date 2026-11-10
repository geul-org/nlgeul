//ff:func feature=token type=reader control=sequence
//ff:what TestLeadText — nothing for whole syllables, 어 for a fused vowel, the consonant otherwise
package token

import "testing"

func TestLeadText(t *testing.T) {
	if leadText(0) != "" || leadText(leadVowel) != "어" || leadText(jongB) != "ㅂ" {
		t.Fatal("leadText")
	}
}
