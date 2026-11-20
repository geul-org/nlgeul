//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what wayBranches — a way as the old structure held it: every choice of one candidate of each of its nodes (the first node's outermost, each in its order) as a branch of their words in order, with the least verdict and the summed preference
package parse

import "math"

func wayBranches(seg Segment) []Branch {
	if len(seg.Nodes) == 0 {
		return nil
	}
	out := []Branch{{Verdict: 1}}
	for _, n := range seg.Nodes {
		var next []Branch
		for _, before := range out {
			for _, b := range n.Contest.Branches {
				words := append(append([]Eojeol(nil), before.Words...), b.Words...)
				next = append(next, Branch{Words: words, Verdict: math.Min(before.Verdict, b.Verdict), Pref: before.Pref + b.Pref})
			}
		}
		out = next
	}
	return out
}
