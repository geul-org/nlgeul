//ff:func feature=rule type=model control=sequence
//ff:what TestShapeValidate — a shape of two runs is valid, one of fewer is not
package rule

import "testing"

func TestShapeValidate(t *testing.T) {
	if (&Shape{Min: 2}).Validate() != nil || (&Shape{Min: 1}).Validate() == nil {
		t.Fatal("validate")
	}
}
