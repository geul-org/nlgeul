//ff:func feature=parse type=helper control=selection
//ff:what writtenNumber — what a text is named when the whole of it is a number that is no run of digits: the plain number name for a run of Sino-Korean numerals in Hangul (사십, 오, 수십) and a run of Han numerals (十, 二三), which the letters then name (numberNameOf: 수:한글십진, 수:한자십진); 수:로마십진 for a Roman numeral written in the Latin capitals I, V and X (II, XIV, V), which its letters do not tell; "" when it is no such number
package parse

func writtenNumber(text string) string {
	switch {
	case sinoTokens(text) > 0, hanRun(text):
		return numName
	case romanWritten(text):
		return romanNumName
	}
	return ""
}
