//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what bestLines — the first line of each stretch of a sentence as the stream sends it first (Next: the first candidate of each node of its best way, Branch.String, joined with " "), in order; a stretch with no reading gives none
package parse

import "strings"

func bestLines(r *Reader, sentence string) []string {
	var out []string
	for _, sg := range streamFirsts(r.Stream(sentence)) {
		var firsts []string
		for _, n := range sg.Nodes {
			firsts = append(firsts, n.Contest.Branches[0].String())
		}
		if len(firsts) > 0 {
			out = append(out, strings.Join(firsts, " "))
		}
	}
	return out
}
