//ff:func feature=rule type=rule control=sequence
//ff:what TestDecimalShape — one dot between numbers has the shape of a decimal; more dots or a leading v do not
package rule

import "testing"

func TestDecimalShape(t *testing.T) {
	one, _ := DecimalShape(markCtx(Mark{Dots: 1}), nil)
	two, _ := DecimalShape(markCtx(Mark{Dots: 2}), nil)
	v, _ := DecimalShape(markCtx(Mark{Dots: 1, VHead: true}), nil)
	if one != -1 || two != 1 || v != 1 {
		t.Fatal(one, two, v)
	}
}
