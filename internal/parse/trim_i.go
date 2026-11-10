//ff:func feature=parse type=reader control=sequence
//ff:what trimI — the copula stem 이 when the tail read it at the end of its stem (것이 → 이); "" when no copula was read there
package parse

import "strings"

func trimI(t tailFact) string {
	if t.copula && strings.HasSuffix(t.stem, "이") {
		return "이"
	}
	return ""
}
