//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what auxNamesOf — the names the auxiliary verbs' stems are read under as auxiliaries (보조 용언:있), in the order of the stems
package parse

func auxNamesOf(aux []string) []string {
	out := make([]string, 0, len(aux))
	for _, w := range aux {
		out = append(out, auxSet+":"+w)
	}
	return out
}
