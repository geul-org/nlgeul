//ff:func feature=token type=reader control=sequence
//ff:what TestRenderHead — a contest of more candidates than head prints the first head and a line of how many were left out, with the indent of a candidate; one of exactly head candidates, or head 0, prints all with no such line; a single candidate is never cut
package token

import (
	"strings"
	"testing"
)

func TestRenderHead(t *testing.T) {
	br := func(kind string) Branch {
		return Branch{Words: []Eojeol{{Parts: []Part{{Text: "곰", Stem: true}, {Text: "이", Name: Name{kind, "이"}}}}}, Verdict: 1, Pref: -2}
	}
	b := Bundle{Text: "곰이", Branches: []Branch{br("주격"), br("보격"), br("서술격"), br("접미"), br("보조사")}}
	got := render(b, 2)
	if len(got) != 4 || !strings.Contains(got[2], "보격") || got[3] != "  (3개 생략됨)" {
		t.Fatal(got)
	}
	if got = render(b, 4); len(got) != 6 || got[5] != "  (1개 생략됨)" {
		t.Fatal(got)
	}
	if got = render(b, 5); len(got) != 6 || strings.Contains(got[5], "생략") {
		t.Fatal(got)
	}
	if got = render(b, 9); len(got) != 6 || strings.Contains(strings.Join(got, "\n"), "생략") {
		t.Fatal(got)
	}
	if got = render(b, 0); len(got) != 6 || strings.Contains(strings.Join(got, "\n"), "생략") {
		t.Fatal(got)
	}
	if got = render(Bundle{Text: "곰이", Branches: b.Branches[:1]}, 1); len(got) != 1 {
		t.Fatal(got)
	}
}
