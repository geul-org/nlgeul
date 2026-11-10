//ff:func feature=mini type=reader control=sequence
//ff:what unfuse — a stem's last syllable without the piece fused into it: the final consonant off, 해 → 하, ㅙ → ㅚ (돼 → 되), 켜 → 키
package mini

func unfuse(r rune) rune {
	r -= rune(jongOf(r))
	if r == '해' {
		return '하'
	}
	if r == '켜' {
		return '키'
	}
	if jungOf(r) == 10 {
		return r + 28
	}
	return r
}
