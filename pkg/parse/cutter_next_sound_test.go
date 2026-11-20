//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCutterNextSound — letters alone in a bracket right after Han characters are told how many characters stand before it (강산 of 江山[강산]: 2, 명월 of 明月(명월): 2); the characters, the brackets, the closing sign with what is written onto it, letters in a bracket with no Han character before it and letters that do not close the bracket are told none
package parse

import (
	"testing"
	"unicode"
)

func TestCutterNextSound(t *testing.T) {
	ix, _ := newIndex(Options{})
	c := newCutter(ix, "江山[강산]에 明月(명월)이 [산] 山[산수 물]", false)
	want := []struct {
		text  string
		gloss int
	}{
		{"江山", 0}, {"[", 0}, {"강산", 2}, {"]에", 0},
		{"明月", 0}, {"(", 0}, {"명월", 2}, {")이", 0},
		{"[", 0}, {"산", 0}, {"]", 0},
		{"山", 0}, {"[", 0}, {"산수", 0}, {"물", 0}, {"]", 0},
	}
	k := 0
	for i := 0; i < len(c.r); {
		if unicode.IsSpace(c.r[i]) {
			i++
			continue
		}
		u, end := c.next(i)
		if k >= len(want) || u.text != want[k].text || u.gloss != want[k].gloss {
			t.Fatalf("unit %d at %d: %q gloss %d", k, i, u.text, u.gloss)
		}
		i, k = end, k+1
	}
	if k != len(want) {
		t.Fatal("units", k)
	}
}
