//ff:func feature=token type=helper control=iteration dimension=1
//ff:what dealsDiff — what differs when up to deals ways are asked of two regions one after another ("" when nothing does): whether there is another, its rank, its words by their spans and its lambda
package token

import "fmt"

func dealsDiff(a, b *region, deals int) string {
	for k := 0; k < deals; k++ {
		sa, oka := a.deal()
		sb, okb := b.deal()
		if oka != okb {
			return fmt.Sprint("deal ", k, oka, okb)
		}
		if !oka {
			return ""
		}
		if sa.Rank != sb.Rank || sa.Lambda != sb.Lambda || segmentLine(sa) != segmentLine(sb) {
			return fmt.Sprint("way ", k, " ", segmentLine(sa), sa.Lambda, " | ", segmentLine(sb), sb.Lambda)
		}
	}
	return ""
}
