//ff:func feature=parse type=helper control=iteration dimension=3
//ff:what streamScore — the verdict and the preference ("%.2f %.2f") of the first candidate of a node of the stretches the stream reads that prints as the given text (Branch.String); "" when no candidate is that text
package parse

import "fmt"

func streamScore(r *Reader, sentence, text string) string {
	for _, rg := range netRegions(r, sentence) {
		for _, n := range rg.nodes {
			for _, b := range n.Contest.Branches {
				if b.String() == text {
					return fmt.Sprintf("%.2f %.2f", b.Verdict, b.Pref)
				}
			}
		}
	}
	return ""
}
