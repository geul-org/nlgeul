//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestBinaryRun — a text of the digits 0 and 1 only is binary, one digit too; any other digit or letter, a sign and no text are not
package token

import "testing"

func TestBinaryRun(t *testing.T) {
	for text, want := range map[string]bool{
		"0": true, "1": true, "10": true, "1010": true, "0001": true,
		"": false, "2": false, "102": false, "1,0": false, "1.01": false, "0x1": false, "-1": false,
	} {
		if got := binaryRun(text); got != want {
			t.Errorf("binaryRun(%q) = %v", text, got)
		}
	}
}
