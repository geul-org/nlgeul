//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what distinct — some values (texts, the candidates of a sign of a pair), each once, in the order they first come; for the candidates of a sign that order is the order of the ways the signs before it may stand open, the first way being the one in which the signs before were read as their first candidates, so a quotation mark written the same at both ends closes first where one of its kind was opened before it, and opens first where none was
package parse

func distinct[T comparable](xs []T) []T {
	var out []T
	seen := map[T]bool{}
	for _, x := range xs {
		if !seen[x] {
			seen[x] = true
			out = append(out, x)
		}
	}
	return out
}
