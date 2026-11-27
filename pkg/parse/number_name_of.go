//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what numberNameOf — what a number is called in the output by the letters it is written in: any digit among them (a sign, a comma or a point may stand beside) 수:십진; a circled number 원문자:수, a parenthesized number 괄호문자:수, one Latin letter 특수:로마자 (letterName); else only Hangul syllables that are a run of numerals (hangulNumeral) 수:한글십진, only Chinese characters 수:한자십진, only Roman numeral letters 수:로마십진; any other letters (Hangul that is no numeral: 몇, a letter numbering a law's items; a circled Latin letter) the plain 수
package parse

import "strings"

func numberNameOf(text string) string {
	if r := []rune(text); len(r) == 1 && letterName(r[0]) != "" {
		return letterName(r[0])
	}
	if strings.ContainsAny(text, "0123456789") {
		return numName + ":" + decimalForm
	}
	form := ""
	for _, r := range text {
		f := numberScript(r)
		if f == "" || (form != "" && f != form) {
			return numName
		}
		form = f
	}
	if form == "" || (form == hangulForm && !hangulNumeral(text)) {
		return numName
	}
	return numName + ":" + form
}
