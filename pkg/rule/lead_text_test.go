//ff:func feature=rule type=helper control=sequence
//ff:what TestLeadText — each final consonant as its letter, a fused vowel as 어, whole syllables as nothing
package rule

import "testing"

func TestLeadText(t *testing.T) {
	if LeadText(JongN) != "ㄴ" || LeadText(JongL) != "ㄹ" || LeadText(JongM) != "ㅁ" || LeadText(JongB) != "ㅂ" || LeadText(JongSS) != "ㅆ" || LeadText(LeadVowel) != "어" || LeadText(0) != "" || LeadText(1) != "" {
		t.Fatal("LeadText")
	}
}
