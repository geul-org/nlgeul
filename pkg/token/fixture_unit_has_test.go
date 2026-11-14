//ff:func feature=token type=helper control=iteration dimension=2
//ff:what unitHas — whether some candidate of some contest of a sentence, as the reader reads it, prints with the given text in it
package token

import "strings"

func unitHas(r *Reader, s, text string) bool {
	for _, c := range r.Contests(s) {
		for _, b := range c.Branches {
			if strings.Contains(b.String(), text) {
				return true
			}
		}
	}
	return false
}
