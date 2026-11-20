//ff:func feature=parse type=helper control=sequence
//ff:what atomNames — the names of an atom joined with ","
package parse

import "strings"

func atomNames(a atom) string {
	return strings.Join(a.names, ",")
}
