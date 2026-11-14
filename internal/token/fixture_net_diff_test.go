//ff:func feature=token type=helper control=iteration dimension=1
//ff:what netDiff — what differs between the stretches of a sentence read by the net and read the old way ("" when nothing does): how many stretches there are, and for the first stretch that differs where (regionDiff, with up to 400 ways asked of each)
package token

import "fmt"

func netDiff(r *Reader, s string) string {
	old, got := oldRegions(r, s), netRegions(r, s)
	if len(old) != len(got) {
		return fmt.Sprint("stretches ", len(old), len(got))
	}
	for i := range old {
		if d := regionDiff(old[i], got[i], 400); d != "" {
			return fmt.Sprint(i, " ", d)
		}
	}
	return ""
}
