//ff:func feature=rule type=model control=sequence
//ff:what TestListSpecName — the name carries the list's name, not its items or its fact
package rule

import "testing"

func TestListSpecName(t *testing.T) {
	a := (&List{Name: "takers", Of: ListAfter, Items: []string{"맺음:니"}}).SpecName()
	b := (&List{Name: "takers", Of: ListForm}).SpecName()
	if a != "list:takers" || b != a {
		t.Fatal(a, b)
	}
}
