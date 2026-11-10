//ff:func feature=token type=helper control=sequence
//ff:what labelOf — the kind before the colon of a marker's name (주격:이 → 주격)
package token

import "strings"

func labelOf(name string) string {
	label, _, _ := strings.Cut(name, ":")
	return label
}
