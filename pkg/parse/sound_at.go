//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what soundAt — how many Han characters stand right before the bracket that opens right before place i of a sentence, when the letters from i close that bracket right at end with its own closing sign (江山[강산]: 2 for 강산, 明月(명월): 2 for 명월); 0 when no such bracket stands there or no Han character stands right before it
package parse

import "unicode"

// The letters alone in such a bracket are the sound of the Han
// characters written right before it: the gloss of 대곤(大棍) the other
// way round. Old writing gives the sound so in a square bracket
// (江山[강산]), and a round one does as well (明月(명월)).
func soundAt(r []rune, i, end int) int {
	if i < 2 || end >= len(r) {
		return 0
	}
	if !(r[i-1] == '[' && r[end] == ']') && !(r[i-1] == '(' && r[end] == ')') {
		return 0
	}
	n := 0
	for j := i - 2; j >= 0 && unicode.Is(unicode.Han, r[j]); j-- {
		n++
	}
	return n
}
