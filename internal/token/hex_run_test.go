//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestHexRun — 0x and one hexadecimal digit or more is a hexadecimal run, whatever the case of the digits and however long; 0x alone, a letter outside a-f, an upper case 0X, a prefix 0b and a text that does not begin with the prefix are not
package token

import "testing"

func TestHexRun(t *testing.T) {
	for label, want := range map[string]bool{
		"0x1234": true, "0xFF": true, "0xaB9": true, "0x0000000000001234": true, "0xf": true,
		"0x": false, "0x12G4": false, "0X1F": false, "0b101": false, "1234": false, "x1F": false, "0x1_2": false, "": false,
	} {
		if got := hexRun(label); got != want {
			t.Errorf("hexRun(%q) = %v", label, got)
		}
	}
}
