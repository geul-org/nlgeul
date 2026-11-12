//ff:func feature=token type=helper control=sequence
//ff:what atomNames — the names of an atom joined with ","
package token

import "strings"

func atomNames(a atom) string {
	return strings.Join(a.names, ",")
}
