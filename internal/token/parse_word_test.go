//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestParseWord — a line gives its key, kind and count; a verb and an adjective are told apart; an unlisted part of speech stands alone; a line with the wrong number of columns, no key or a bad count is refused
package token

import "testing"

func TestParseWord(t *testing.T) {
	if key, kind, count, err := parseWord("마시\t동사\t규칙\t900"); err != nil || key != "마시" || kind != kindVerb || count != 900 {
		t.Fatal(key, kind, count, err)
	}
	if _, kind, _, _ := parseWord("작\t형용사\t규칙\t7"); kind != kindAdj {
		t.Fatal(kind)
	}
	if _, kind, _, err := parseWord("of\t전치사\t\t3"); err != nil || kind != kindAlone {
		t.Fatal(kind, err)
	}
	for _, bad := range []string{"마시\t동사\t900", "\t동사\t규칙\t900", "마시\t동사\t규칙\t많이", "마시\t동사\t규칙\t-1"} {
		if _, _, _, err := parseWord(bad); err == nil {
			t.Errorf("%q accepted", bad)
		}
	}
}
