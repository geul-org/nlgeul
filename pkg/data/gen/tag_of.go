//ff:func feature=data type=helper control=sequence
//ff:what tagOf — the tag of a compatibility decomposition (<circle>, <compat> ...), or empty when there is none
package main

import "strings"

func tagOf(decomp string) string {
	if !strings.HasPrefix(decomp, "<") {
		return ""
	}
	return strings.Fields(decomp)[0]
}
