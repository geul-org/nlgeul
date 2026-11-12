//ff:func feature=token type=helper control=iteration dimension=1
//ff:what crossParts — the readings given, each once for each of the parts, with its word k holding that one part
package token

func crossParts(branches []Branch, k int, parts []Part) []Branch {
	var out []Branch
	for _, cur := range branches {
		out = append(out, readParts(cur, k, parts)...)
	}
	return out
}
