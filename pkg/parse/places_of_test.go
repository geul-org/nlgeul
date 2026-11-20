//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPlacesOf — no table when no unit holds a space; with spaces inside a span of code the table has one place more than the letters of the units, a space keeps the place before it and the letters after it go on without it
package parse

import "testing"

func TestPlacesOf(t *testing.T) {
	if got := placesOf([]unit{{text: "가나"}, {text: "`ab`을"}}); got != nil {
		t.Fatal(got)
	}
	if got := placesOf(nil); got != nil {
		t.Fatal(got)
	}
	got := placesOf([]unit{{text: "`a b`을"}, {text: "가  나"}})
	want := []int{0, 1, 2, 2, 3, 4, 5, 6, 6, 6, 7}
	if len(got) != len(want) {
		t.Fatal(got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatal(got, want)
		}
	}
}
