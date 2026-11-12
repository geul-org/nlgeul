//ff:func feature=token type=helper control=iteration dimension=1
//ff:what markNames — the names of some readings of a sign, in order
package token

func markNames(reads []markRead) []string {
	out := make([]string, len(reads))
	for i, r := range reads {
		out[i] = r.name
	}
	return out
}
