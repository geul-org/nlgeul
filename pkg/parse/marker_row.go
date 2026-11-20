//ff:func feature=parse type=reader control=sequence
//ff:what markerRow — one closing of this layer's own table (gated-closings.tsv) read from a line of it (form, what must stand before it, kinds, ways of attaching)
package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
)

func markerRow(row data.Row) (Marker, error) {
	return MarkerOf(row, 0)
}
