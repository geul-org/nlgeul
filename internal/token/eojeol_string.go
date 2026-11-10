//ff:func feature=token type=helper control=iteration dimension=1
//ff:what String — a word printed: its parts joined with "+"
package token

import "strings"

func (e Eojeol) String() string {
	parts := make([]string, 0, len(e.Parts))
	for _, p := range e.Parts {
		parts = append(parts, p.String())
	}
	return strings.Join(parts, "+")
}
