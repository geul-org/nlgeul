//ff:func feature=token type=helper control=iteration dimension=1
//ff:what waysBracket — the brackets of a contest of ways drawn: the inner bracket of each way (tagged with its rank and lambda) and the outer bracket of the contest (tagged "어절경합 [시작-끝]" with the place of the stretch); lines is the number of lines of all the ways
package token

import "fmt"

func waysBracket(ways []boxWay, segs []Segment, lines int) *bracket {
	b := newBracket(lines)
	starts := wayStarts(ways)
	for i, w := range ways {
		b.drawInner(starts[i], starts[i]+w.last, wayTag(segs[i]))
	}
	b.drawOuter(starts, fmt.Sprintf("어절경합 [%d-%d]", segs[0].Region.Start, segs[0].Region.End))
	return b
}
