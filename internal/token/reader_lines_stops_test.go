//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderLinesStops — with no word table, the signs at the edge of a word come up as lines of their own: commas between words are 쉼표, 밥 먹었어? is 물음표, 좋다! 느낌표, 뭐... 줄임표, 좋아~ 물결표, a lone 쌍점 after a word 쌍점, and a sentence's last 마침표; the signs in a row are one stretch (??, !!, ~~~)
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesStops(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"사과, 포도, 수박을 샀다", []string{"사과", ",[쉼표]", "포도", ",[쉼표]", "수박을", "샀다"}},
		{"밥 먹었어?", []string{"밥", "먹었어", "?[물음표]"}},
		{"정말?? 좋다!!", []string{"정말", "??[물음표]", "좋다", "!![느낌표]"}},
		{"좋다!", []string{"좋다", "![느낌표]"}},
		{"뭐...", []string{"뭐", "...[줄임표]"}},
		{"좋아~ 정말~~~", []string{"좋아", "~[물결표]", "정말", "~~~[물결표]"}},
		{"제목: 내용", []string{"제목", ":[쌍점]", "내용"}},
		{"했다.", []string{"했다", ".[마침표]"}},
	} {
		lines := r.Lines(c.in, false)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
	}
}
