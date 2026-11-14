//ff:func feature=token type=helper control=sequence
//ff:what TestUnspeltAt — a fused ㅆ is read under 되 with the 어 left out (됬), not under another syllable nor where no ㅆ sits, whatever follows; a fused 어 after 되 only when more of the tail follows or the word may end there, never after another syllable; another lead never
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestUnspeltAt(t *testing.T) {
	if !unspeltAt('됬', rule.JongSS, false) || !unspeltAt('됬', rule.JongSS, true) {
		t.Fatal("ㅆ under 되")
	}
	if unspeltAt('갔', rule.JongSS, false) || unspeltAt('되', rule.JongSS, false) || unspeltAt('됐', rule.JongSS, false) {
		t.Fatal("ㅆ under another syllable, or no ㅆ under 되")
	}
	if !unspeltAt('되', rule.LeadVowel, true) {
		t.Fatal("어 after 되 with more of the tail")
	}
	if unspeltAt('되', rule.LeadVowel, false) || unspeltAt('가', rule.LeadVowel, true) || unspeltAt('돼', rule.LeadVowel, true) {
		t.Fatal("어 after 되 with nothing open, or after another syllable")
	}
	if unspeltAt('되', rule.JongN, true) || unspeltAt('됬', 0, true) || unspeltAt('된', rule.JongN, true) {
		t.Fatal("another lead")
	}
}
