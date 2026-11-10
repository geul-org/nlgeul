//ff:func feature=token type=reader control=sequence
//ff:what TestCore — punctuation around a field is dropped; inside it stays
package token

import "testing"

func TestCore(t *testing.T) {
	if string(core("\"먹자.\"")) != "먹자" || string(core("3.5를,")) != "3.5를" || string(core("...")) != "" {
		t.Fatal("core")
	}
}
