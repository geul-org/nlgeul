//ff:func feature=token type=helper control=iteration dimension=1
//ff:what textsOf — the letters of each bundle, in order
package token

func textsOf(bundles []Bundle) []string {
	var out []string
	for _, b := range bundles {
		out = append(out, b.Text)
	}
	return out
}
