//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamApartAngle — with no word table, a < and a > each written with a space on both sides are read as a pair when both stand so in the sentence (< 제목 >: 열기:1 and 닫기:1, each beside its letter with no name), and what they enclose is one deeper; a < or a > written apart with no partner written apart (a comparison) and one next to a hyphen stay their letters
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamApartAngle(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in    string
		named int
		want  []string
	}{
		{"오늘의 소식 < 한빛상사 > 크게 올랐다", 2, []string{"오늘", "소식", "<[열기:1]\t성립 1.00", "<\t", "한빛상사", ">[닫기:1]\t성립 1.00", ">\t", "크게", "올랐다"}},
		{"< 알림(둘) > 보라", 2, []string{"<[열기:1]", "알림", "([열기:2]", "둘", ")[닫기:2]", ">[닫기:1]", "보라"}},
		{"값이 3 < 5 이면 된다", 0, []string{"값이", "3", "<\t", "5", "이면", "된다"}},
		{"조건은 값 > 15 일 때다", 0, []string{"조건은", "값", ">\t", "15", "일", "때다"}},
		{"<이름 > 보라", 1, []string{"<[열기:1]", "이름", ">\t", "보라"}},
	} {
		lines := scoredLines(r, c.in)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
		text := strings.Join(lines, "\n")
		if got := strings.Count(text, "<[열기:") + strings.Count(text, ">[닫기:"); got != c.named {
			t.Errorf("%q: %d named angle brackets, want %d, in\n%s", c.in, got, c.named, text)
		}
	}
}
