//ff:func feature=rule type=model control=sequence
//ff:what TestShapeSpecName — the name carries the reading
package rule

import "testing"

func TestShapeSpecName(t *testing.T) {
	if got := (&Shape{Of: "날짜:-"}).SpecName(); got != "shape:날짜:-" {
		t.Fatal(got)
	}
}
