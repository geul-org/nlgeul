//ff:func feature=token type=helper control=iteration dimension=1
//ff:what ordered — the candidates of a sign of a pair, each once, in the order they first come: the ways the signs before it may stand open give them in turn, and the first way is the one in which the signs before were read as their first candidates, so a quotation mark written the same at both ends closes first where one of its kind was opened before it, and opens first where none was
package token

func ordered(candidates []pairing) []pairing {
	var out []pairing
	seen := map[pairing]bool{}
	for _, rd := range candidates {
		if !seen[rd] {
			seen[rd] = true
			out = append(out, rd)
		}
	}
	return out
}
