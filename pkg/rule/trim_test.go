//ff:func feature=rule type=helper control=sequence
//ff:what TestTrim — punctuation and symbols around a field are dropped; inside it they stay
package rule

import "testing"

func TestTrim(t *testing.T) {
	if string(Trim("\"먹자.\"")) != "먹자" || string(Trim("3.5를,")) != "3.5를" || string(Trim("...")) != "" {
		t.Fatal("Trim")
	}
}
