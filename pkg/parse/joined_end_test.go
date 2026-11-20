//ff:func feature=parse type=helper control=sequence
//ff:what TestJoinedEnd — Hangul after a hyphen or an underscore at a path's end is a part of its name; Hangul right after a Latin letter or a slash is not, nor is anything when nothing follows the letter
package parse

import "testing"

func TestJoinedEnd(t *testing.T) {
	if got := joinedEnd([]rune("~/a/b-가격"), 6); got != 8 {
		t.Errorf("hyphen: %d", got)
	}
	if got := joinedEnd([]rune("~/a/b_가격"), 6); got != 8 {
		t.Errorf("underscore: %d", got)
	}
	if got := joinedEnd([]rune("~/a/b를"), 5); got != 5 {
		t.Errorf("particle: %d", got)
	}
	if got := joinedEnd([]rune("~/a/가격"), 4); got != 4 {
		t.Errorf("slash: %d", got)
	}
	if got := joinedEnd([]rune("~/a/b-"), 6); got != 6 {
		t.Errorf("nothing after: %d", got)
	}
	if got := joinedEnd([]rune("가격"), 0); got != 0 {
		t.Errorf("no plain letter: %d", got)
	}
}
