//ff:func feature=parse type=helper control=selection
//ff:what writtenNumber — what a text is named when the whole of it is a number that is no run of digits: the plain number name for a run of Sino-Korean numerals in Hangul (사십, 오, 수십), which the letters then name (numberNameOf: 수:한글십진); "" when it is no such number
package parse

func writtenNumber(text string) string {
	switch {
	case sinoTokens(text) > 0:
		return numName
	}
	return ""
}
