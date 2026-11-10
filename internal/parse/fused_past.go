//ff:func feature=parse type=reader control=sequence
//ff:what fusedPast — is r a stem vowel fused with 었/았 (갔, 섰, 봤, 줬, 졌, 냈, 셌, 됐, 했): final ㅆ on one of the vowels of endings.md section 4
package parse

func fusedPast(r rune) bool {
	v := jungOf(r)
	return jongOf(r) == jongSS && (v == 0 || v == 1 || v == 4 || v == 5 || v == 6 || v == 9 || v == 10 || v == 14)
}
