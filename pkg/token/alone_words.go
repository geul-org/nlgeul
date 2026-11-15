//ff:func feature=token type=helper control=iteration dimension=1
//ff:what aloneWords — the words of the closed sets a word alone in a pair of brackets is not read as (aloneBarred: the bound nouns and the determiners of the grammar tables), whose letters may also be read as another word there
package token

import "github.com/geul-org/nlgeul/pkg/parse"

func aloneWords() []string {
	sets := parse.ClosedWords()
	var out []string
	for _, s := range aloneBarred {
		out = append(out, sets[s]...)
	}
	return out
}
