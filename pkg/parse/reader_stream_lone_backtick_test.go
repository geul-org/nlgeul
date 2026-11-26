//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamLoneBacktick — with no word table, a backtick that no other backtick of the sentence pairs with into a span of code opens a pair whose closing sign stands after the sentence (열기:1), beside its letter with no name: before a word, alone at the sentence's end, and after a span of code; a pair opened after it is one deeper; backticks two by two are still a span of code and open nothing
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamLoneBacktick(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in    string
		opens int
		want  []string
	}{
		{"`기쁨은 앎에서 온다.", 1, []string{"`[열기:1]\t성립 1.00", "`\t", "기쁨", "앎", "온다", ".[마침표]"}},
		{"아래 목록에서 `", 1, []string{"아래", "목록", "`[열기:1]", "`\t"}},
		{"거기서 열린 `웃어라! (하나)", 1, []string{"거기서", "열린", "`[열기:1]", "`\t", "웃어라", "![느낌표]", "([열기:2]", "하나", ")[닫기:2]"}},
		{"`a b`를 쓰고 `끝", 1, []string{"`a b`[코드]+를[목적격:를]", "쓰고", "`[열기:1]", "끝"}},
		{"`a b`를 쓰고 `c d`도 쓴다", 0, []string{"`a b`[코드]+를[목적격:를]", "쓰고", "`c d`[코드]+도[보조사:도]", "쓴다"}},
	} {
		lines := scoredLines(r, c.in)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
		if got := strings.Count(strings.Join(lines, "\n"), "`[열기:"); got != c.opens {
			t.Errorf("%q: %d opening backticks, want %d, in\n%s", c.in, got, c.opens, strings.Join(lines, "\n"))
		}
	}
}
