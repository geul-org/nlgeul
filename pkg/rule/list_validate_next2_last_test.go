//ff:func feature=rule type=model control=sequence
//ff:what TestListValidateNext2Last — ListNext2Last is the last fact a list may be asked for: a list of it with a name is valid, one past it or one with no name is refused
package rule

import "testing"

func TestListValidateNext2Last(t *testing.T) {
	if err := (&List{Name: "lasts", Of: ListNext2Last}).Validate(); err != nil {
		t.Error(err)
	}
	if err := (&List{Name: "lasts", Of: ListNext2Last + 1}).Validate(); err == nil {
		t.Error("a fact past the last is refused")
	}
	if err := (&List{Of: ListNext2Last}).Validate(); err == nil {
		t.Error("a list with no name is refused")
	}
}
