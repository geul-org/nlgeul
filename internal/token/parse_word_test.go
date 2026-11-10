//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestParseWord — a line gives its key, kind and count; an unlisted part of speech stands alone; a line with the wrong number of columns, no key or a bad count is refused
package token

import "testing"

func TestParseWord(t *testing.T) {
	if key, w, err := parseWord("마시\t동사\t규칙\t900"); err != nil || key != "마시" || w != (word{kindVerbal, 900}) {
		t.Fatal(key, w, err)
	}
	if _, w, err := parseWord("of\t전치사\t\t3"); err != nil || w.kinds != kindAlone {
		t.Fatal(w, err)
	}
	for _, bad := range []string{"마시\t동사\t900", "\t동사\t규칙\t900", "마시\t동사\t규칙\t많이", "마시\t동사\t규칙\t-1"} {
		if _, _, err := parseWord(bad); err == nil {
			t.Errorf("%q accepted", bad)
		}
	}
}
