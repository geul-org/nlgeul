//ff:func feature=token type=reader control=sequence
//ff:what markerRow — one closing of this layer's own list read from a line of its table (form, what must stand before it, kinds, ways of attaching)
package token

import (
	"github.com/geul-org/nlgeul/data"
	"github.com/geul-org/nlgeul/pkg/parse"
)

func markerRow(row data.Row) (parse.Marker, error) {
	return parse.MarkerOf(row, 0)
}
