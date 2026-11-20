//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what pieceNamed — the piece among some that is the candidate of the given name (the empty piece when there is none)
package parse

func pieceNamed(pieces map[string]piece, name string) piece {
	for _, p := range pieces {
		if p.names == name {
			return p
		}
	}
	return piece{}
}
