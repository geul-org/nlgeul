//ff:func feature=token type=helper control=iteration dimension=1
//ff:what distinct — some texts, each once, in the order they first come
package token

func distinct(texts []string) []string {
	var out []string
	seen := map[string]bool{}
	for _, t := range texts {
		if !seen[t] {
			seen[t] = true
			out = append(out, t)
		}
	}
	return out
}
