//ff:func feature=token type=helper control=iteration dimension=1
//ff:what readParts — a candidate given once for each of the parts, with its word k holding that one part (withPart)
package token

func readParts(br Branch, k int, parts []Part) []Branch {
	var out []Branch
	for _, p := range parts {
		out = append(out, withPart(br, k, p))
	}
	return out
}
