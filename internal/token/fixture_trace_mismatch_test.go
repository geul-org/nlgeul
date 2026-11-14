//ff:func feature=token type=helper control=iteration dimension=1
//ff:what traceMismatch — what differs between the trace lines the net and the old reading give for each unit of a sentence ("" when nothing does): the unit, how many lines, or the first line that differs
package token

import "fmt"

func traceMismatch(r *Reader, s string) string {
	for k, u := range unitsOf(r.ix, s, true) {
		_, want := readUnit(r.ix, u, true)
		_, got := readUnitNet(r.ix, u, true)
		if len(got) != len(want) {
			return fmt.Sprint("unit ", k, " lines ", len(got), len(want))
		}
		if d := traceDiff(got, want); d != "" {
			return fmt.Sprint("unit ", k, " ", d)
		}
	}
	return ""
}
