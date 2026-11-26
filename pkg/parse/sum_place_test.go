//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestSumPlace — one hyphen alone stands where a sum is written when a bracket stands open before it in the first way, or when a number (or a number and its unit) ends right before the space before it and a digit stands right after the space after it; not a row of hyphens, not a hyphen written onto a word, not at the sentence's edge, not between words outside a bracket, not inside a quotation mark alone
package parse

import "testing"

func TestSumPlace(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text    string
		i, end  int
		open    string
		want    bool
		because string
	}{
		{"100 - 20", 4, 5, "", true, "between two numbers"},
		{"3,000원 - 1,000원", 7, 8, "", true, "after a number and its unit"},
		{"나 - 다", 2, 3, "[", true, "inside a bracket that stands open"},
		{"나 - 다", 2, 3, "\"", false, "a quotation mark is no bracket of a sum"},
		{"나 - 다", 2, 3, "", false, "between words"},
		{"100 - 가", 4, 5, "", false, "no digit after it"},
		{"가 - 20", 2, 3, "", false, "no number before it"},
		{"100 -- 20", 4, 6, "", false, "a row of hyphens"},
		{"100 -20", 4, 5, "", false, "written onto the number after it"},
		{"100- 20", 3, 4, "", false, "written onto the number before it"},
		{"- 20", 0, 1, "", false, "at the sentence's start"},
		{"100 -", 4, 5, "", false, "at the sentence's end"},
		{"100 - ", 4, 5, "", false, "nothing after the space"},
		{"100 ~ 20", 4, 5, "", false, "another sign"},
	} {
		cut := newCutter(ix, c.text, false)
		cut.open = []string{c.open}
		if got := cut.sumPlace(c.i, c.end); got != c.want {
			t.Errorf("%q (%s): %v", c.text, c.because, got)
		}
	}
	none := newCutter(ix, "가 - 나", false)
	none.open = nil
	if none.sumPlace(2, 3) {
		t.Fatal("no way of standing open")
	}
}
