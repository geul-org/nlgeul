//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what againKnown — the ways of a stretch after the first, asked one at a time with AgainHead and printed, until one of known words only or none is left (then "none" is printed)
package parse_test

import (
	"fmt"

	"github.com/geul-org/nlgeul/pkg/parse"
)

func againKnown(st *parse.Stream, region parse.Span, head int) {
	for {
		seg, ok := st.AgainHead(region, head)
		if !ok {
			fmt.Println("  none of known words only")
			return
		}
		printWay(seg)
		if wayKnown(seg) {
			return
		}
	}
}
