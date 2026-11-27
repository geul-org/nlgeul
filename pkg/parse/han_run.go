//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what hanRun — is a text a run of Han numerals (hanNumerals: 十, 二三, 二〇二四) and nothing else; false for an empty text
package parse

import "slices"

func hanRun(text string) bool {
	for _, r := range text {
		if !slices.Contains(hanNumerals, string(r)) {
			return false
		}
	}
	return text != ""
}
