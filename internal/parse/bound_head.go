//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what boundHead — does the word begin with a dependent noun and hold nothing else but its particles or a tail on it (것, 때까지, 거지, 것이다)
package parse

import (
	"slices"
	"strings"
)

func boundHead(w *word) bool {
	host, _, _ := strings.Cut(w.chain.marked, "+")
	if slices.Contains(boundNouns, host) || slices.Contains(boundNouns, w.text) {
		return true
	}
	for _, t := range w.tails {
		stem := strings.TrimSuffix(t.stem, trimI(t))
		if strings.HasPrefix(w.text, t.stem) && slices.Contains(boundNouns, stem) {
			return true
		}
	}
	return false
}
