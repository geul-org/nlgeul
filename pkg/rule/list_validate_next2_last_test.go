//ff:func feature=rule type=model control=sequence
//ff:what TestListValidateNext2Last — a list of ListNext2Last with a name is valid, one with no name is refused, and one past the last fact (ListNextName) is refused
package rule

import "testing"

func TestListValidateNext2Last(t *testing.T) {
	if err := (&List{Name: "lasts", Of: ListNext2Last}).Validate(); err != nil {
		t.Error(err)
	}
	if err := (&List{Name: "lasts", Of: ListNextName + 1}).Validate(); err == nil {
		t.Error("a fact past the last is refused")
	}
	if err := (&List{Of: ListNext2Last}).Validate(); err == nil {
		t.Error("a list with no name is refused")
	}
}
