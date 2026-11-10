//ff:func feature=token type=helper control=iteration dimension=1
//ff:what String — a word printed: its parts joined with "+"; a word of markers alone that attaches to the unit before begins with "+"
package token

import "strings"

func (e Eojeol) String() string {
	parts := make([]string, 0, len(e.Parts))
	for _, p := range e.Parts {
		parts = append(parts, p.String())
	}
	if e.Attach {
		return "+" + strings.Join(parts, "+")
	}
	return strings.Join(parts, "+")
}
