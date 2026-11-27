//ff:func feature=parse type=helper control=sequence
//ff:what bracketedAt — do the letters from place i to end of a sentence stand alone in a round bracket that is written right onto a word (크로뮴(III): the bracket opens right before i, right after a Hangul syllable, a Latin letter or a digit, and closes right at end)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func bracketedAt(r []rune, i, end int) bool {
	if i < 2 || end >= len(r) || r[i-1] != '(' || r[end] != ')' {
		return false
	}
	return rule.JongOf(r[i-2]) >= 0 || rule.RunRune(r[i-2])
}
