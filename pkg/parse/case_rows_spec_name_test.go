//ff:func feature=parse type=model control=sequence
//ff:what TestCaseRowsSpecName — the criterion's name is caseRows, whatever rows it holds
package parse

import "testing"

func TestCaseRowsSpecName(t *testing.T) {
	if got := (&caseRows{Rows: []caseException{{Word: "담소"}}}).SpecName(); got != "caseRows" {
		t.Fatal(got)
	}
}
