//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestMarksOf — a word of a stem alone (no markers) gives none, one marker gives its name, and several give each name apart and in order
package rule

import (
	"slices"
	"testing"
)

func TestMarksOf(t *testing.T) {
	cases := []struct {
		marks string
		want  []string
	}{
		{"", nil},
		{"부사격:에서", []string{"부사격:에서"}},
		{"부사격:에서+보조사:는", []string{"부사격:에서", "보조사:는"}},
		{"맺음:다+보조사:는+보조사:도", []string{"맺음:다", "보조사:는", "보조사:도"}},
	}
	for _, c := range cases {
		got := MarksOf(Near{Stem: "학교", Marks: c.marks})
		if !slices.Equal(got, c.want) || (c.want == nil) != (got == nil) {
			t.Errorf("MarksOf(%q) = %q; want %q", c.marks, got, c.want)
		}
	}
	if got := MarksOf(Near{}); got != nil {
		t.Errorf("empty Near: %q", got)
	}
}
