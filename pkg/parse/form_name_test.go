//ff:func feature=parse type=helper control=sequence
//ff:what TestFormName — a form is called by its shape and what it writes, a colon between them
package parse

import "testing"

func TestFormName(t *testing.T) {
	if got := (form{'❶', "검은 원", "수", "1", "1", ""}).name(); got != "검은 원:수" {
		t.Fatal(got)
	}
}
