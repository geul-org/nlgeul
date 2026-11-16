//ff:func feature=token type=reader control=sequence
//ff:what markerRow — one closing of this layer's own table (gated-closings.tsv) read from a line of it (form, what must stand before it, kinds, ways of attaching)
package token

import (
	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/parse"
)

func markerRow(row data.Row) (parse.Marker, error) {
	return parse.MarkerOf(row, 0)
}
