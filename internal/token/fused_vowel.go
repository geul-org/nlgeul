//ff:func feature=token type=reader control=sequence
//ff:what fusedVowel — is r a stem vowel that 어/아 fuses into: ㅘ ㅝ ㅕ ㅙ and 해 always; ㅏ ㅓ ㅐ ㅔ (no visible change) only when more of the tail follows
package token

func fusedVowel(r rune, more bool) bool {
	if jongOf(r) != 0 {
		return false
	}
	v := int(r-0xAC00) / 28 % 21
	seen := v == 6 || v == 9 || v == 10 || v == 14
	unseen := v == 0 || v == 1 || v == 4 || v == 5
	return r == '해' || seen || (more && unseen)
}
