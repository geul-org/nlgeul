//ff:func feature=token type=reader control=iteration dimension=1
//ff:what glossAt — how many Han characters a round bracket that opens at place i of a sentence holds, when it holds nothing but Han characters and is closed (대곤(大棍): 2 at the bracket); 0 when no such bracket stands there
package token

import "unicode"

// The bracket is the round one of 「한글 맞춤법」 부록 문장 부호
// 소괄호 (2): it shows the original writing beside the Korean one.
func glossAt(r []rune, i int) int {
	if i >= len(r) || r[i] != '(' {
		return 0
	}
	for j := i + 1; j < len(r); j++ {
		if r[j] == ')' {
			return j - i - 1
		}
		if !unicode.Is(unicode.Han, r[j]) {
			return 0
		}
	}
	return 0
}
