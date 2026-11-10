//ff:func feature=token type=helper control=iteration dimension=1
//ff:what stretchEnd — where the letters from a place of a sentence stop: at the next space, at the next sign of a pair, or at the end
package token

import "unicode"

func stretchEnd(r []rune, from int, pairs map[int]int) int {
	for from < len(r) && !unicode.IsSpace(r[from]) {
		if _, paired := pairs[from]; paired {
			break
		}
		from++
	}
	return from
}
