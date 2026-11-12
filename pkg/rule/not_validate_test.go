//ff:func feature=rule type=model control=sequence
//ff:what TestNotValidate — the criterion is always valid
package rule

import "testing"

func TestNotValidate(t *testing.T) {
	if err := (&Not{}).Validate(); err != nil {
		t.Fatal(err)
	}
}
