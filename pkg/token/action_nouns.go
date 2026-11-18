//ff:func feature=token type=helper control=iteration dimension=1
//ff:what actionNouns — the nouns of a word table that name an action or an event: a nominal key of two syllables or more whose letters with 하 are a verb key of the table (발병 → 발병하다), sorted
package token

import (
	"slices"
	"unicode/utf8"
)

// A noun that 하다 makes a verb of names an action or an event (school
// grammar: 서술성 명사). A noun of one syllable is left out: its 하 verb
// is mostly another word (정하다, 권하다), and surnames are of one.
func actionNouns(lex map[string]word) []string {
	var out []string
	for k, w := range lex {
		if w.kinds&kindNominal == 0 || utf8.RuneCountInString(k) < 2 {
			continue
		}
		if lex[k+haStem].kinds&kindVerb != 0 {
			out = append(out, k)
		}
	}
	slices.Sort(out)
	return out
}
