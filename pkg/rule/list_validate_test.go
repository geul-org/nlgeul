//ff:func feature=rule type=model control=iteration dimension=1
//ff:what TestListValidate — a list with a name and a fact from ListForm to ListClosing is valid, even with no items; one with no name, or a fact out of that range, is refused
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
		{List{Name: "forms", Of: ListClosing + 1}, false},
	}
	for _, c := range cases {
		if err := c.l.Validate(); (err == nil) != c.ok {
			t.Errorf("Validate(%+v) = %v", c.l, err)
		}
	}
}
