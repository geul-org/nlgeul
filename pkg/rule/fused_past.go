//ff:func feature=rule type=helper control=sequence
//ff:what FusedPast — 도움 함수(규칙 아님). 줄기 끝 모음에 었/았이 합쳐진 음절인가(갔, 섰, 봤, 줬, 졌, 냈, 셌, 됐, 했): 받침 ㅆ이고 모음이 ㅏ ㅐ ㅓ ㅔ ㅕ ㅘ ㅙ ㅝ 가운데 하나
package rule

func FusedPast(r rune) bool {
	v := JungOf(r)
	return JongOf(r) == JongSS && (v == 0 || v == 1 || v == 4 || v == 5 || v == 6 || v == 9 || v == 10 || v == 14)
}
