//ff:func feature=mini type=rule control=sequence
//ff:what TestCore — quotation marks and sentence punctuation around the word are dropped; inside it they stay
package mini

import "testing"

func TestCore(t *testing.T) {
	if got := string(core("\"먹자.\"")); got != "먹자" {
		t.Fatal(got)
	}
	if got := string(core("3.5를,")); got != "3.5를" {
		t.Fatal(got)
	}
	if got := string(core("...")); got != "" {
		t.Fatal(got)
	}
}
