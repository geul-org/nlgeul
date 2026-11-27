//ff:func feature=rule type=model control=iteration dimension=1
//ff:what TestListValidate — a list with a name and a fact from ListForm to ListNext2Last (the last fact, the name of the last piece of the second word after a candidate in its sentence) is valid, even with no items; one with no name, or a fact out of that range, is refused
package rule

import "testing"

func TestListValidate(t *testing.T) {
	cases := []struct {
		l  List
		ok bool
	}{
		{List{Name: "forms", Of: ListForm, Items: []string{"듯하"}}, true},
		{List{Name: "closings", Of: ListClosing}, true},
		{List{Name: "words", Of: ListWord, Items: []string{}}, true},
		{List{Of: ListForm, Items: []string{"듯하"}}, false},
		{List{Name: "forms"}, false},
		{List{Name: "forms", Of: ListForm - 1}, false},
		{List{Name: "selves", Of: ListSelfStem}, true},
		{List{Name: "nexts", Of: ListNextStem, Items: []string{"되"}}, true},
		{List{Name: "marks", Of: ListSelfMarks, Items: []string{"부사격:에서"}}, true},
		{List{Name: "after", Of: ListNext2Stem, Items: []string{"되"}}, true},
		{List{Name: "before", Of: ListPrevKind, Items: []string{"닫기"}}, true},
		{List{Name: "stems", Of: ListPrevStem, Items: []string{"관하"}}, true},
		{List{Name: "after", Of: ListNextKind, Items: []string{"닫기"}}, true},
		{List{Name: "letters", Of: ListSelfText, Items: []string{"한"}}, true},
		{List{Name: "names", Of: ListPrevName, Items: []string{"열쇠:동사"}}, true},
		{List{Name: "lasts", Of: ListNext2Last, Items: []string{"꾸밈:는"}}, true},
		{List{Name: "forms", Of: ListNextName + 1}, false},
	}
	for _, c := range cases {
		if err := c.l.Validate(); (err == nil) != c.ok {
			t.Errorf("Validate(%+v) = %v", c.l, err)
		}
	}
}
