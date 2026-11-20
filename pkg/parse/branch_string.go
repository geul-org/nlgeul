//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what String — a branch printed: its words with a space between them
package parse

import "strings"

func (b Branch) String() string {
	words := make([]string, 0, len(b.Words))
	for _, w := range b.Words {
		words = append(words, w.String())
	}
	return strings.Join(words, " ")
}
