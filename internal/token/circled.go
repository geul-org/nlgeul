//ff:func feature=token type=helper control=selection
//ff:what circled — the number of a way in a circle: ①~⑳, ㉑~㉟, ㊱~㊿; a number above 50 (or below 1) in round brackets, (51)
package token

import "fmt"

func circled(n int) string {
	switch {
	case n >= 1 && n <= 20:
		return string(rune(0x2460 + n - 1))
	case n >= 21 && n <= 35:
		return string(rune(0x3251 + n - 21))
	case n >= 36 && n <= 50:
		return string(rune(0x32B1 + n - 36))
	}
	return fmt.Sprintf("(%d)", n)
}
