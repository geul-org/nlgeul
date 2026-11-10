//ff:func feature=rule type=rule control=sequence
//ff:what TestMarkAlone — applies to a single dot
package rule

import "testing"

func TestMarkAlone(t *testing.T) {
	one, _ := MarkAlone(markCtx(Mark{Run: 1}), nil)
	three, _ := MarkAlone(markCtx(Mark{Run: 3}), nil)
	if one != 1 || three != -1 {
		t.Fatal(one, three)
	}
}
