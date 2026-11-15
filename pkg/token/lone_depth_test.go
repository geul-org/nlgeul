//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestLoneDepth — a closing bracket with none of its kind open closes the outermost (loneBracketDepth, 1) whichever bracket it is, a quotation mark with none of its kind open 0
package token

import "testing"

func TestLoneDepth(t *testing.T) {
	if loneBracketDepth != 1 {
		t.Fatal(loneBracketDepth)
	}
	for sign, want := range map[rune]int{')': 1, '}': 1, ']': 1, '>': 1, '」': 1, '』': 1, '》': 1, '〉': 1, '"': 0, '\'': 0, '”': 0, '’': 0} {
		if got := loneDepth(sign); got != want {
			t.Errorf("%c: %d, want %d", sign, got, want)
		}
	}
}
