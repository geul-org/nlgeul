//ff:func feature=token type=reader control=sequence
//ff:what TestUnitsOf — a sentence is cut at spaces and at the signs of a pair: an opening bracket alone, a closing one with the particle written onto it, a pair that spans a space, a span of code with its spaces as one unit; a sign with no partner stays in its unit; an empty sentence has no units
package token

import (
	"fmt"
	"testing"
)

func TestUnitsOf(t *testing.T) {
	got := fmt.Sprint(unitsOf("곰(큰 곰)이 `a b`를 본다"))
	want := fmt.Sprint([]unit{{"곰", 0, ""}, {"(", 1, "여는 묶음표"}, {"큰", 0, ""}, {"곰", 0, ""}, {")이", 1, "닫는 묶음표"}, {"`a b`를", 5, codeName}, {"본다", 0, ""}})
	if got != want {
		t.Fatal(got)
	}
	if got = fmt.Sprint(unitsOf("곰(이 \"간다")); got != fmt.Sprint([]unit{{"곰(이", 0, ""}, {"\"간다", 0, ""}}) {
		t.Fatal("signs with no partner", got)
	}
	if unitsOf("  ") != nil {
		t.Fatal("no units")
	}
}
