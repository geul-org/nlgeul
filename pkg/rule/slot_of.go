//ff:func feature=rule type=helper control=selection
//ff:what SlotOf — 도움 함수(규칙 아님). 음절이 선어말 칸이면 그 칸(시 1, 었·았·였 2, 겠 3, 더 4), 아니면 0
package rule

func SlotOf(r rune) int {
	switch r {
	case '시':
		return SlotHonor
	case '었', '았', '였':
		return SlotPast
	case '겠':
		return SlotGuess
	case '더':
		return SlotRecall
	}
	return 0
}
