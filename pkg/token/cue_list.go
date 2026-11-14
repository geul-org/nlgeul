//ff:func feature=token type=reader control=iteration dimension=1
//ff:what cueList — the words of a list as cues at place p, when the rule of that article is on
package token

func (sc *scan) cueList(p int, article string, words []string) {
	if !sc.ix.spacing[article] {
		return
	}
	for _, w := range words {
		sc.cueWord(p, w)
	}
}
