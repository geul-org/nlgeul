//ff:func feature=token type=helper control=iteration dimension=2
//ff:what TestStateKeyOfWord — of all the candidates of the old reading of every made-up sentence (a word table, spacing on, no table) those whose last words have the same key have the same word (parts and all), which is what lets the net keep one state for them
package token

import "testing"

func TestStateKeyOfWord(t *testing.T) {
	for name, r := range netReaders(t) {
		for _, s := range netSentences() {
			if d := keyedWordsDiffer(r, s); d != "" {
				t.Fatal(name, s, d)
			}
		}
	}
}
