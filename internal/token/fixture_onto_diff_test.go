//ff:func feature=token type=helper control=iteration dimension=1
//ff:what ontoDiff — what differs between the stretches of a sentence read by the net and read the old way, but for the numbers of birth, which the net counts from another place ("" when nothing does): the stretches, then for the first that differs its span, its nodes (their candidates in order), the best and rest of each, the words that follow each, how many ways there are, and up to 400 ways asked of both
package token

import (
	"fmt"
	"reflect"
)

func ontoDiff(r *Reader, s string) string {
	old, got := oldRegions(r, s), netRegions(r, s)
	if len(old) != len(got) {
		return fmt.Sprint("stretches ", len(old), len(got))
	}
	for i := range old {
		a, b := old[i], got[i]
		if a.span != b.span || len(a.nodes) != len(b.nodes) {
			return fmt.Sprint(i, " span or nodes ", a.span, b.span, len(a.nodes), len(b.nodes))
		}
		if d := nodesDiff(a, b); d != "" {
			return fmt.Sprint(i, " ", d)
		}
		if !reflect.DeepEqual(a.best, b.best) || !reflect.DeepEqual(a.rest, b.rest) || !reflect.DeepEqual(a.next, b.next) || a.ways.Cmp(b.ways) != 0 {
			return fmt.Sprint(i, " best, rest, next or ways ", a.ways, b.ways)
		}
		if d := regionRest(a, b, 400); d != "" {
			return fmt.Sprint(i, " ", d)
		}
	}
	return ""
}
