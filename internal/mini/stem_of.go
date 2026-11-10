//ff:func feature=mini type=reader control=sequence
//ff:what stemOf — the stem as text, with a fused piece taken out of its last syllable: ㅆ off (갔 → 가), 해 → 하, ㅙ → ㅚ (돼 → 되)
package mini

// stemOf undoes only the fusions endings.md section 4 lists with one
// source (하 → 해, ㅚ → ㅙ); other fused vowels stay as they are.
func stemOf(s []rune) string {
	out := append([]rune(nil), s...)
	last := &out[len(out)-1]
	if fusedPast(*last) {
		*last -= jongSS
	}
	if *last == '해' {
		*last = '하'
	}
	if jongOf(*last) == 0 && jungOf(*last) == 10 {
		*last += 28
	}
	return string(out)
}
