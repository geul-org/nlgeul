//ff:func feature=token type=helper control=sequence
//ff:what windowText — a made-up sentence of the word table words-far.tsv written with no spaces, repeated, so that it is one long unit whose words are plain
package token

import "strings"

func windowText(times int) string {
	return strings.Repeat("고양이가마당에서물을마신다", times)
}
