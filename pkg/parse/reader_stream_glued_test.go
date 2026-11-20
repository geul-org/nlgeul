//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamGlued — Glued tells that no space stood before a contest: false for the first, true for the second contest of a unit, true for a unit that follows one with no space after it (a plus, a sign of a pair), false after a space
package parse

import (
	"reflect"
	"testing"
)

func TestReaderStreamGlued(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for text, want := range map[string][]bool{
		"가 나":     {false, false},
		"곰+돌":     {false, true, true},
		"(가)를 나":  {false, true, true, true, false},
		"\"6\" 나": {false, true, true, false},
		"가\n나":    {false, false},
	} {
		var got []bool
		for _, b := range streamContests(r, text) {
			got = append(got, b.Glued)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%q: %v, want %v", text, got, want)
		}
	}
}
