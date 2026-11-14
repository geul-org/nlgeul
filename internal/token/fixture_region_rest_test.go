//ff:func feature=token type=helper control=sequence
//ff:what regionRest — what differs in the ways two regions hand out ("" when nothing does): up to deals ways asked of both, and last how many walks were skipped
package token

import "fmt"

func regionRest(a, b *region, deals int) string {
	if d := dealsDiff(a, b, deals); d != "" {
		return d
	}
	if a.skipped != b.skipped {
		return fmt.Sprint("skipped ", a.skipped, b.skipped)
	}
	return ""
}
