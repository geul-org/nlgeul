//ff:func feature=token type=reader control=iteration dimension=1
//ff:what Cuts — for measuring: one line for each unit, its text, the word ends (token offsets) of each reading that is awake, best first, readings separated by ";", and after another tab the same for the readings that are asleep
package token

import "strings"

func (r *Reader) Cuts(sentence string) []string {
	var out []string
	for _, field := range strings.Fields(sentence) {
		sc := &scan{ix: r.ix, s: trimUnit(field)}
		awake, asleep := wake(sc.search())
		out = append(out, string(sc.s)+"\t"+endsOf(awake)+"\t"+endsOf(asleep))
	}
	return out
}
