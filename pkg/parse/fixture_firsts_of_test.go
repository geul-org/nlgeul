//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what firstsOf — the first line of a way (the first candidate of each of its nodes, Branch.String, joined with " ") when the letters of its nodes are the given word; none otherwise
package parse

import "strings"

func firstsOf(seg Segment, word string) []string {
	var text string
	var firsts []string
	for _, n := range seg.Nodes {
		text += n.Contest.Text
		firsts = append(firsts, n.Contest.Branches[0].String())
	}
	if text != word || len(firsts) == 0 {
		return nil
	}
	return []string{strings.Join(firsts, " ")}
}
