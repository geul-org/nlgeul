//ff:func feature=rule type=helper control=sequence
//ff:what FusedVowel — 도움 함수(규칙 아님). 줄기 끝 모음에 어/아가 합쳐진 음절인가: ㅘ ㅝ ㅕ ㅙ와 해는 늘, ㅏ ㅓ ㅐ ㅔ(겉으로 안 바뀜)는 뒤에 꼬리가 더 있을 때만
package rule

func FusedVowel(r rune, more bool) bool {
	if JongOf(r) != 0 {
		return false
	}
	v := JungOf(r)
	seen := v == 6 || v == 9 || v == 10 || v == 14
	unseen := v == 0 || v == 1 || v == 4 || v == 5
	return r == '해' || seen || (more && unseen)
}
