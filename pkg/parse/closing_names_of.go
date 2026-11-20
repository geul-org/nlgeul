//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what closingNamesOf — the names the index gives the entries of the closings of the closing table whose kinds hold the kind given (rule.KindLink: 연결 and 종결|연결), in the order of the table
package parse

func closingNamesOf(kind int) []string {
	var out []string
	for _, m := range ClosingMarkers() {
		if m.Kind&kind == 0 {
			continue
		}
		for _, e := range closingEntries(m) {
			out = append(out, e.name)
		}
	}
	return out
}
