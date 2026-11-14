//ff:func feature=token type=helper control=sequence
//ff:what boxTitle — the title of a node's box: its place, its letters and how many candidates its contest has (all of them, not only those shown)
package token

import "fmt"

func boxTitle(n Node) string {
	return fmt.Sprintf("[%d-%d] %s ─ 후보경합 %d", n.Span.Start, n.Span.End, n.Contest.Text, len(n.Contest.Branches))
}
