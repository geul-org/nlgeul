//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamLoneBracket — with no word table, a closing bracket written onto a word with none of its kind open in the sentence closes as 닫기:1 (its opening sign stood before the sentence), and so does one right after a mark that ends a sentence (다.』, 오?」), with a comma, a full stop or a particle after it read as before (+이라는); what stands open is left as it is, so a pair opened after it still closes at its own depth, and one of another kind opened before it stays open; a named sign written right after a closing bracket (ㆍ) is a sign of its own; a round bracket right after a digit with none of its kind open numbers an item of a list: it closes nothing and is read as 닫기:0 beside its letter with no name, with a particle after it read as before and a pair opened after it at its own depth; any other bracket right after a digit, one with a space before it, and one whose pair is open give no lone closing
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamLoneBracket(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		lone bool
		want []string
	}{
		{"판결 참조), 피고가 이겼다.", true, []string{"판결", "참조", ")[닫기:1]\t성립 1.00", ",[쉼표]", "피고가"}},
		{"할 것이다}.", true, []string{"것이다", "}[닫기:1]\t성립 1.00", ".[마침표]"}},
		{"처신 논란>이라는 제목", true, []string{"논란", ">[닫기:1]\t성립 1.00", "+이[서술격:이]+라는[꾸밈:라는]\t성립 1.00", "제목"}},
		{"참조) 등을 (보라) 했다", true, []string{"참조", ")[닫기:1]", "등을", "([열기:1]", "보라", ")[닫기:1]", "했다"}},
		{"“여당 논란>이라는 제목", true, []string{"“[열기:1]", "여당", "논란", ">[닫기:1]", "+이[서술격:이]+라는[꾸밈:라는]", "제목"}},
		{"그래서 싸운다.』", true, []string{"그래서", "싸운다", ".[마침표]", "』[닫기:1]\t성립 1.00"}},
		{"무슨 일이오?」", true, []string{"무슨", "일이오", "?[물음표]", "」[닫기:1]\t성립 1.00"}},
		{"(책) 참조", false, []string{"([열기:1]", "책", ")[닫기:1]", "참조"}},
		{"그는 1) 갔다", true, []string{"그는", "1", ")[닫기:0]\t성립 1.00", ")\t", "갔다"}},
		{"비고 3)에 적은 값 (참고) 포함", true, []string{"비고", "3", ")[닫기:0]", ")\t", "+에[부사격:에]", "적은", "값", "([열기:1]", "참고", ")[닫기:1]", "포함"}},
		{"표(주 2)와 같다", false, []string{"표", "([열기:1]", "주", "2", ")[닫기:1]", "+와[부사격:와]", "같다"}},
		{"크기는 3] 이다", false, []string{"크기는", "3", "]\t", "이다"}},
		{"그는 ) 갔다", false, []string{"그는", ")\t", "갔다"}},
		{"성명(이름)ㆍ주소", false, []string{"성명", "([열기:1]", "이름", ")[닫기:1]", "ㆍ[가운뎃점]", "주소"}},
	} {
		lines := scoredLines(r, c.in)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
		if got := strings.Count(strings.Join(lines, "\n"), "[닫기:"); !c.lone && got > strings.Count(c.in, "(") {
			t.Errorf("%q: a closing bracket closes with none of its kind open, in\n%s", c.in, strings.Join(lines, "\n"))
		}
	}
}
