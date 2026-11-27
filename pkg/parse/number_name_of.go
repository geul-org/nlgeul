//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what numberNameOf — what a number is called in the output by the letters it is written in: any digit among them (a sign, a comma or a point may stand beside) 수:십진; else only Hangul syllables 수:한글십진, only Chinese characters 수:한자십진, only Roman numeral letters 수:로마십진; any other letters (a circled or parenthesized number, one Latin letter) the plain 수
package parse

import "strings"

func numberNameOf(text string) string {
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
	if form == "" {
		return numName
	}
	return numName + ":" + form
}
