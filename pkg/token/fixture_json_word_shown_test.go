//ff:func feature=token type=helper control=iteration dimension=2
//ff:what jsonWordShown — the candidates of a word set down as data, best first, each as its parts joined with "+", a part as its letters with its name in brackets (the letters alone for a part with no name)
package token

import "strings"

func jsonWordShown(w JSONWord) []string {
	var out []string
	for _, c := range w.Candidates {
		parts := make([]string, 0, len(c.Parts))
		for _, p := range c.Parts {
			if p.Name == nil {
				parts = append(parts, p.Text)
				continue
			}
			parts = append(parts, p.Text+"["+*p.Name+"]")
		}
		out = append(out, strings.Join(parts, "+"))
	}
	return out
}
