//ff:func feature=token type=helper control=sequence
//ff:what TestUnspeltAt — a fused ㅆ is read under 되 with the 어 left out (됬), not under another syllable nor where no ㅆ sits, whatever follows; a fused 어 after 되 only when more of the tail follows or the word may end there, never after another syllable; another lead never
package token

import "testing"

func TestUnspeltAt(t *testing.T) {
	if !unspeltAt('됬', jongSS, false) || !unspeltAt('됬', jongSS, true) {
		t.Fatal("ㅆ under 되")
	}
	if unspeltAt('갔', jongSS, false) || unspeltAt('되', jongSS, false) || unspeltAt('됐', jongSS, false) {
		t.Fatal("ㅆ under another syllable, or no ㅆ under 되")
	}
	if !unspeltAt('되', leadVowel, true) {
		t.Fatal("어 after 되 with more of the tail")
	}
	if unspeltAt('되', leadVowel, false) || unspeltAt('가', leadVowel, true) || unspeltAt('돼', leadVowel, true) {
		t.Fatal("어 after 되 with nothing open, or after another syllable")
	}
	if unspeltAt('되', jongN, true) || unspeltAt('됬', 0, true) || unspeltAt('된', jongN, true) {
		t.Fatal("another lead")
	}
}
