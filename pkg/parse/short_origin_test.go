//ff:func feature=parse type=helper control=sequence
//ff:what TestShortOrigin — 지는 and 고는 are what a shortened closing of the table stands for (진, 곤); 어서는 and 을지는 are not
package parse

import "testing"

func TestShortOrigin(t *testing.T) {
	if !shortOrigin("지는") || !shortOrigin("고는") {
		t.Fatal("지는, 고는")
	}
	if shortOrigin("어서는") || shortOrigin("을지는") || shortOrigin("") {
		t.Fatal("어서는, 을지는")
	}
}
