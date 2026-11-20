//ff:func feature=parse type=helper control=sequence
//ff:what labelOf — the kind before the colon of a name (주격 of 주격:이, 의존명사 of 의존명사:수); the name itself when it has no colon
package parse

import "strings"

func labelOf(name string) string {
	label, _, _ := strings.Cut(name, ":")
	return label
}
