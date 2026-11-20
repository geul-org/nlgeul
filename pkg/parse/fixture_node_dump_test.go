//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what nodeDump — a node in lines: its span and letters, then each candidate in order with its lambda and preference
package parse

import "fmt"

func nodeDump(n Node) []string {
	out := []string{fmt.Sprintf("  node %v %s", n.Span, n.Contest.Text)}
	for _, br := range n.Contest.Branches {
		out = append(out, fmt.Sprintf("    %s %v %v", br.String(), br.Lambda(), br.Pref))
	}
	return out
}
