//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wordOf — one word of a reading: the stem from b (a closed-set word or a new word) and the markers after it, printed with "+" and the graph names, with its verdict and preference
package token

import (
	"math"
	"strings"
)

// A stem syllable costs log10(1/syllables) of preference whether the
// stem is a key or not: the tables give keys no counts.
func (sc *scan) wordOf(b, end int, t tail) eojeol {
	stem := string(sc.s[b:t.start])
	e := eojeol{end: end, marked: stem, names: newWord, verdict: sc.ix.unknown}
	if k, ok := sc.ix.words[stem]; ok {
		e.names, e.verdict, e.marked = k.name, 1, stem+"["+k.name+"]"
	}
	e.pref = float64(t.start-b) * math.Log10(1.0/syllables)
	names := []string{e.names}
	for _, p := range t.pieces {
		e.marked += "+" + p.text + "[" + p.names + "]"
		e.verdict, e.pref = math.Min(e.verdict, p.verdict), e.pref+p.pref
		names = append(names, p.names)
	}
	e.names = strings.Join(names, " + ")
	return e
}
