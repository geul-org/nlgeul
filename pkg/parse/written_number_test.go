//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestWrittenNumber — a text that is a number and no run of digits is named: by the plain number name where its letters tell the form, by the name of the way it is written where they do not; a run of digits, a native numeral and another word have no such name
package parse

import "testing"

func TestWrittenNumber(t *testing.T) {
	for text, want := range map[string]string{"사십": numName, "오": numName, "수십": numName, "삼십만": numName, "十": numName, "二三": numName,
		"40": "", "세": "", "육칠": "", "곰": "", "CD": "", "ii": "", "": ""} {
		if got := writtenNumber(text); got != want {
			t.Errorf("writtenNumber(%q) = %q", text, got)
		}
	}
}
