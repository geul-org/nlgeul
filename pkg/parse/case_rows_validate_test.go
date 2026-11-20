//ff:func feature=parse type=model control=sequence
//ff:what TestCaseRowsValidate — a criterion that holds a row is valid, one with none is an error
package parse

import "testing"

func TestCaseRowsValidate(t *testing.T) {
	if err := (&caseRows{Rows: []caseException{{Word: "담소"}}}).Validate(); err != nil {
		t.Fatal(err)
	}
	if err := (&caseRows{}).Validate(); err == nil {
		t.Fatal("no row accepted")
	}
}
