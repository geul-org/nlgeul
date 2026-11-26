//ff:func feature=parse type=rule control=sequence
//ff:what signGraphOf — the judgment graph of one candidate of one named sign: that of the candidate's name (signGraph), but for a sign written for a dash that is drawn with too (plainDashes: the hyphen, the box-drawing line) the graph that reads its place (dashGraph); the number of edges is returned
package parse

import (
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func signGraphOf(sign rune, name string) (*toulmin.Graph, int) {
	if strings.ContainsRune(plainDashes, sign) {
		return dashGraph(sign, name)
	}
	return signGraph(name)
}
