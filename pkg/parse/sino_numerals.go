//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what sinoNumerals — of the numbers before a unit that are no run of digits, those a unit of sinoUnits (인) is read after: the runs of Sino-Korean numerals (삼인) and of Han numerals (三인), not a native numeral (네, 한) or a Latin letter; for any other unit all of them
package parse

import "slices"

func sinoNumerals(form string, words []numeral) []numeral {
	if !slices.Contains(sinoUnits, form) {
		return words
	}
	var out []numeral
	for _, w := range words {
		if sinoTokens(w.text) > 0 || hanRun(w.text) {
			out = append(out, w)
		}
	}
	return out
}
