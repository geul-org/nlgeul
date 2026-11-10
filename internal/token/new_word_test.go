//ff:func feature=token type=rule control=sequence
//ff:what TestNewWord — a new word has its low verdict, and none when it stands after a boundary and reads as markers
package token

import "testing"

func TestNewWord(t *testing.T) {
	ix := scanOf("x").ix
	if ix.newWord(false) != ix.unknown || ix.newWord(true) > 0 {
		t.Fatal(ix.newWord(false), ix.newWord(true))
	}
}
