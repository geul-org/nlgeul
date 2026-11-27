//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what hangulNumeral — is a text a run of the numerals written in Hangul (hangulNumerals: 사십, 열두, 스물한, 두세, 육칠만) and nothing else, 수 counting at its front before a greater number (수십, 수백만); false for an empty text and for any other Hangul (몇, 여러, 가)
package parse

import "strings"

func hangulNumeral(text string) bool {
	if rest, ok := strings.CutPrefix(text, string(sinoSome)); ok && rest != "" && strings.ContainsRune(sinoBigs, []rune(rest)[0]) {
		text = rest
	}
	if text == "" {
		return false
	}
	end := make([]bool, len(text)+1)
	end[0] = true
	for i := range text {
		for _, n := range numeralLengths(text[i:], end[i]) {
			end[i+n] = true
		}
	}
	return end[len(text)]
}
