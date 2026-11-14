//ff:func feature=token type=helper control=iteration dimension=2
//ff:what scoreFaults — the branches of the contests of more than one word that have no scores of their own kept, or whose scores do not add up to the branch's preference (sum) and verdict (least); each as its printed line
package token

import "math"

func scoreFaults(contests []Contest) []string {
	var bad []string
	for _, c := range contests {
		for _, b := range c.Branches {
			sum, least := 0.0, 1.0
			for _, s := range scoresOf(b) {
				sum, least = sum+s.pref, math.Min(least, s.verdict)
			}
			if (len(b.Words) > 1 && len(b.scores) != len(b.Words)) || math.Abs(sum-b.Pref) > 1e-9 || least != b.Verdict {
				bad = append(bad, b.String())
			}
		}
	}
	return bad
}
