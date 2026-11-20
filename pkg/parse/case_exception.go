//ff:type feature=parse type=model
//ff:what caseException — one row of the case exception table: the graph it attacks (its name in the index), the word that graph must not read (the unit's letters up to the piece's end), the reading the row keeps ("" for none), the kind of its ground and the ground, and the line it is read from
package parse

import "github.com/geul-org/nlgeul/pkg/data"

// Row is kept for the line a trace names and for an error about the
// row's graph, which is found only when the index is built.
type caseException struct {
	Graph      string
	Word       string
	Reading    string
	GroundKind string
	Ground     string
	Row        data.Row
}
