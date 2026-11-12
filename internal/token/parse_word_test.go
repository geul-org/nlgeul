//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestParseWord — a line gives its key, kind, irregular type and count; a verb and an adjective are told apart; an unlisted part of speech stands alone; a line of five columns gives its right form (an empty fifth column or four columns give ""); a line with other than four or five columns, no key or a bad count is refused
package token

import "testing"

func TestParseWord(t *testing.T) {
	if key, kind, irr, count, _, err := parseWord("마시\t동사\t규칙\t900"); err != nil || key != "마시" || kind != kindVerb || irr != 0 || count != 900 {
		t.Fatal(key, kind, irr, count, err)
	}
	if _, kind, irr, _, _, _ := parseWord("덥\t형용사\tㅂ\t7"); kind != kindAdj || irr != irrB {
		t.Fatal(kind, irr)
	}
	if _, kind, _, _, _, err := parseWord("of\t전치사\t\t3"); err != nil || kind != kindAlone {
		t.Fatal(kind, err)
	}
	if key, kind, irr, count, right, err := parseWord("맞기\t동사\t규칙\t0\t맡기"); err != nil || key != "맞기" || kind != kindVerb || irr != 0 || count != 0 || right != "맡기" {
		t.Fatal(key, kind, irr, count, right, err)
	}
	if _, _, _, _, right, err := parseWord("마시\t동사\t규칙\t900\t"); err != nil || right != "" {
		t.Fatal("an empty fifth column", right, err)
	}
	if _, _, _, _, right, err := parseWord("마시\t동사\t규칙\t900"); err != nil || right != "" {
		t.Fatal("four columns", right, err)
	}
	for _, bad := range []string{"마시\t동사\t900", "맞기\t동사\t규칙\t0\t맡기\t여섯", "\t동사\t규칙\t900", "마시\t동사\t규칙\t많이", "마시\t동사\t규칙\t-1"} {
		if _, _, _, _, _, err := parseWord(bad); err == nil {
			t.Errorf("%q accepted", bad)
		}
	}
}
