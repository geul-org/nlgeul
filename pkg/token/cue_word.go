//ff:func feature=token type=reader control=sequence
//ff:what cueWord — a word w that stands in the letters from p, when it does not end inside an atom: it may stand as a word of no markers where it is (stands, by where it ends and begins), and when the letter before it is a Hangul syllable the unit's letters before it are a word that may end there too (shows)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Where a sign or a script change already shows a boundary (a letter
// before that is not a Hangul syllable) nothing more is shown.
func (sc *scan) cueWord(p int, w string) {
	q := matchAt(sc.s, p, w)
	if q < 0 || sc.inside[q] {
		return
	}
	sc.stand(q, p)
	if p > 0 && rule.JongOf(sc.s[p-1]) >= 0 {
		if sc.shows[p] == nil {
			sc.shows[p] = map[int]bool{}
		}
		sc.shows[p][q] = true
	}
}
