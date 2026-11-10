//ff:func feature=token type=reader control=iteration dimension=1
//ff:what stretch — one reading's way through the stretch (from, to]: its words there, with the least verdict and the summed preference
package token

import "math"

func stretch(p *path, from, to int) Branch {
	b := Branch{Verdict: 1}
	for _, e := range wordsOf(p) {
		if e.end > from && e.end <= to {
			b.Words = append(b.Words, e.word)
			b.Verdict, b.Pref = math.Min(b.Verdict, e.verdict), b.Pref+e.pref
		}
	}
	return b
}
