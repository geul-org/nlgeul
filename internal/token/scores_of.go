//ff:func feature=token type=helper control=sequence
//ff:what scoresOf — the verdict and the preference of each word of a branch: those kept with it, and for a branch made without them the branch's own values for each word (right for a branch of one word) with no number of birth (noBorn)
package token

import "slices"

func scoresOf(br Branch) []wordScore {
	if len(br.scores) == len(br.Words) {
		return br.scores
	}
	return slices.Repeat([]wordScore{{verdict: br.Verdict, pref: br.Pref, born: noBorn}}, len(br.Words))
}
