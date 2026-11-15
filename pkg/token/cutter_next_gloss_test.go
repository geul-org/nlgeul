//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCutterNextGloss — each unit is told how many Han characters a round bracket right after it holds alone: plain letters cut right before one (대곤 before (大棍): 2) and a span of code before one (`a` before (大): 1); the bracket, the characters, a closing sign with what is written onto it, letters with a space before the bracket, a sign that cuts the sentence and the letters after it are told none
package token

import (
	"testing"
	"unicode"
)

func TestCutterNextGloss(t *testing.T) {
	ix, _ := newIndex(Options{})
	c := newCutter(ix, "대곤(大棍)을 관치 (官治) `a`(大) 곰+돌", false)
	want := []struct {
		text  string
		gloss int
	}{
		{"대곤", 2}, {"(", 0}, {"大棍", 0}, {")을", 0},
		{"관치", 0}, {"(", 0}, {"官治", 0}, {")", 0},
		{"`a`", 1}, {"(", 0}, {"大", 0}, {")", 0},
		{"곰", 0}, {"+", 0}, {"돌", 0},
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
