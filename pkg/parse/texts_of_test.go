//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what textsOf — the letters of each contest, in order
package parse

func textsOf(contests []Contest) []string {
	var out []string
	for _, b := range contests {
		out = append(out, b.Text)
	}
	return out
}
