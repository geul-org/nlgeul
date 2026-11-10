//ff:func feature=rule type=rule control=sequence
//ff:what TestMarkRepeated — applies to dots in a row
package rule

import "testing"

func TestMarkRepeated(t *testing.T) {
	one, _ := MarkRepeated(markCtx(Mark{Run: 1}), nil)
	three, _ := MarkRepeated(markCtx(Mark{Run: 3}), nil)
	if one != -1 || three != 1 {
		t.Fatal(one, three)
	}
}
