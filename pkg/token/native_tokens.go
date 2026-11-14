//ff:func feature=token type=helper control=selection
//ff:what nativeTokens — how many native numerals a text is made of, in the form before a unit: one of the ones, alone or after 열; 열, 스무, 몇, 여러 alone; 0 when it is none of these
package token

import (
	"slices"
	"strings"
)

func nativeTokens(text string) int {
	switch rest, teen := strings.CutPrefix(text, "열"); {
	case slices.Contains(nativeAlone, text) || slices.Contains(nativeOnes, text):
		return 1
	case teen && slices.Contains(nativeOnes, rest):
		return 2
	}
	return 0
}
