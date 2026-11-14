//ff:func feature=token type=helper control=iteration dimension=2
//ff:what boxWaysLines — the ways of a stretch cut more than one way (the first ones, as many as segs) each drawn as boxes at the width of the widest box of all of them, one way after the other, with the brackets at the right; a line that counts the ways left out follows the last box
package token

import (
	"math/big"
	"strings"
)

func boxWaysLines(segs []Segment, head int) []string {
	ways := make([]boxWay, len(segs))
	for i, seg := range segs {
		ways[i] = newBoxWay(seg, head)
	}
	width := widestWay(ways)
	var lines []string
	for i := range ways {
		ways[i].draw(width)
		lines = append(lines, ways[i].lines...)
	}
	br := waysBracket(ways, segs, len(lines))
	for i, l := range lines {
		lines[i] = strings.TrimRight(l+br.text(i), " ")
	}
	if left := new(big.Int).Sub(segs[0].Ways, big.NewInt(int64(len(segs)))); left.Sign() > 0 {
		lines = append(lines, omitLine(width, left))
	}
	return lines
}
