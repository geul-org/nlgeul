//ff:func feature=parse type=reader control=sequence
//ff:what unfuse — a stem's last syllable without the piece fused into it: the final consonant off, and the vowel back as endings.md section 4 gives it (해 → 하, ㅘ → ㅗ, ㅝ → ㅜ, ㅕ → ㅣ, ㅙ → ㅚ)
package parse

func unfuse(r rune) rune {
	r -= rune(jongOf(r))
	if r == '해' {
		return '하'
	}
	back := map[int]int{6: 20, 9: 8, 10: 11, 14: 13}
	if to, ok := back[jungOf(r)]; ok {
		return r + rune((to-jungOf(r))*28)
	}
	return r
}
