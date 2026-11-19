//ff:func feature=rule type=model control=sequence
//ff:what TestListValidatePrevName — a list of ListPrevName with a name is valid, one with no name is refused
package rule

import "testing"

func TestListValidatePrevName(t *testing.T) {
	if err := (&List{Name: "names", Of: ListPrevName}).Validate(); err != nil {
		t.Error(err)
	}
	if err := (&List{Of: ListPrevName}).Validate(); err == nil {
		t.Error("a list with no name is refused")
	}
}
