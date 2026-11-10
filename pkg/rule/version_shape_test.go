//ff:func feature=rule type=rule control=sequence
//ff:what TestVersionShape — v1.2, 1.2.3 and 3.14 have the shape of a version (a number with one dot is a decimal and a version); a.b does not
package rule

import "testing"

func TestVersionShape(t *testing.T) {
	v, _ := VersionShape(markCtx(Mark{Numeric: true, VHead: true, Dots: 1}), nil)
	three, _ := VersionShape(markCtx(Mark{Numeric: true, Dots: 2}), nil)
	decimal, _ := VersionShape(markCtx(Mark{Numeric: true, Dots: 1}), nil)
	name, _ := VersionShape(markCtx(Mark{Dots: 2}), nil)
	if v != -1 || three != -1 || decimal != -1 || name != 1 {
		t.Fatal(v, three, decimal, name)
	}
}
