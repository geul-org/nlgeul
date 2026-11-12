//ff:func feature=token type=reader control=sequence
//ff:what TestCutterNextCut — a plus between letters is a unit of its own (no head, the trace lines of the graphs on it) and the letters on both sides are units; a quote right after a digit is a unit with its head marked as a possible unit of measure, and the ways the signs may stand open keep the ways before it; another quote is not marked
package token

import (
	"reflect"
	"testing"
)

func TestCutterNextCut(t *testing.T) {
	ix, _ := newIndex(Options{})
	c := newCutter(ix, "곰+돌 7\" 다\"", true)
	plus, end := c.next(1)
	if plus.text != "+" || plus.head != 0 || plus.name != "" || end != 2 || len(plus.notes) != 2 || plus.measure || c.open[0] != "" {
		t.Fatal("the plus", plus, end)
	}
	if left, end := c.next(0); left.text != "곰" || end != 1 {
		t.Fatal("before the plus", left, end)
	}
	if right, end := c.next(2); right.text != "돌" || !right.spaced || end != 3 {
		t.Fatal("after the plus", right, end)
	}
	seven, _ := c.next(4)
	quote, end := c.next(5)
	if seven.text != "7" || quote.text != "\"" || !quote.measure || quote.head != 1 || !reflect.DeepEqual(quote.reads, []pairing{{openName, 1}}) || end != 6 || !reflect.DeepEqual(c.open, []string{"\"", ""}) {
		t.Fatal("the quote after a digit", quote, end, c.open)
	}
	c.next(7)
	later, _ := c.next(8)
	if later.measure || later.text != "\"" {
		t.Fatal("a quote after a letter", later)
	}
}
