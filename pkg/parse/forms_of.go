//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what formsOf — the different ways of cutting, each written by its spans (lineKey)
package parse

func formsOf(cuts []cut) map[string]bool {
	out := map[string]bool{}
	for _, c := range cuts {
		out[lineKey(c.spans)] = true
	}
	return out
}
