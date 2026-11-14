//ff:func feature=token type=helper control=iteration dimension=1
//ff:what traceDiff — the first trace line that differs between two lists of the same length, "" when none does
package token

func traceDiff(a, b []string) string {
	for i := range a {
		if i < len(b) && a[i] != b[i] {
			return a[i] + " | " + b[i]
		}
	}
	return ""
}
