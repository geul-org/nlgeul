//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReadUnitNetTrace — the trace lines of every unit of every made-up sentence read by the net are those of the old reading, in the same order, with each reader
package token

import "testing"

func TestReadUnitNetTrace(t *testing.T) {
	for name, r := range netReaders(t) {
		for _, s := range netSentences() {
			if d := traceMismatch(r, s); d != "" {
				t.Fatal(name, s, d)
			}
		}
	}
}
