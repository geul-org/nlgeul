//ff:func feature=token type=helper control=iteration dimension=1
//ff:what oneWordEach — whether every candidate of a node holds exactly one word
package token

func oneWordEach(n Node) bool {
	for _, b := range n.Contest.Branches {
		if len(b.Words) != 1 {
			return false
		}
	}
	return true
}
