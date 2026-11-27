//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestRomanWritten — a Roman numeral written in the Latin capitals I, V and X in the numeral's own order is one (I, II, IV, IX, XIV, XXIII, XXXIX); letters out of that order, the numerals of L, C, D and M, small letters and an empty text are not
package parse

import "testing"

func TestRomanWritten(t *testing.T) {
	for text, want := range map[string]bool{"I": true, "V": true, "X": true, "II": true, "III": true, "IV": true, "VI": true, "IX": true, "XIV": true, "XX": true, "XXIII": true, "XXXIX": true,
		"IIII": false, "VV": false, "IIV": false, "VX": false, "XXXX": false, "IXI": false, "CD": false, "MD": false, "XL": false, "MIX": false, "ii": false, "iv": false, "X1": false, "": false} {
		if got := romanWritten(text); got != want {
			t.Errorf("romanWritten(%q) = %v", text, got)
		}
	}
}
