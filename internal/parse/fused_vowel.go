//ff:func feature=parse type=reader control=sequence
//ff:what fusedVowel — is p a stem vowel fused with 어/아: ㅘ ㅝ ㅕ ㅙ and 해 always; ㅏ ㅓ ㅐ ㅔ (no visible change) only when more of the tail follows
package parse

func fusedVowel(p rune, more bool) bool {
	if jongOf(p) != 0 {
		return false
	}
	v := jungOf(p)
	seen := v == 6 || v == 9 || v == 10 || v == 14
	unseen := v == 0 || v == 1 || v == 4 || v == 5
	return p == '해' || seen || (more && unseen)
}
