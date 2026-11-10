//ff:func feature=parse type=helper control=sequence
//ff:what suffixAt — where form starts if s[:end] ends with it
package parse

func suffixAt(s []rune, end int, form string) (int, bool) {
	start := end - len([]rune(form))
	if start < 0 {
		return 0, false
	}
	return start, string(s[start:end]) == form
}
