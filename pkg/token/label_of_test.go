//ff:func feature=token type=helper control=sequence
//ff:what TestLabelOf — the kind before the colon
package token

import "testing"

func TestLabelOf(t *testing.T) {
	if labelOf("주격:이") != "주격" || labelOf("없음") != "없음" {
		t.Fatal("labelOf")
	}
}
