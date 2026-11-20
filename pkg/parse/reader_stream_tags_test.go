//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamTags — with no word table, # and @ are signs of their own: #태그 is 샵 or 태그, @홍길동에게 골뱅이 or 명칭; # 제목1 at the start of a line is 제목1 or 샵, ## 제목2 is 제목2 or the letters (##, ###, #### the same with their counts); a # in the middle of a line, alone, is 샵, ## there has no name, and C# has a # with no name
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamTags(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"#태그 @홍길동에게", []string{"#[샵]", "#[태그]", "태그", "@[골뱅이]", "@[명칭]", "홍길동에게"}},
		{"# 제목1", []string{"#[제목1]", "#[샵]", "제목1"}},
		{"## 제목2", []string{"##[제목2]", "##\t", "제목2"}},
		{"### 제목3", []string{"제목3"}},
		{"#### 제목4", []string{"제목4"}},
		{"가 # 나", []string{"가", "#[샵]", "나"}},
		{"가 ## 나", []string{"가", "##", "나"}},
		{"C#", []string{"C[열:영숫자]", "#"}},
	} {
		lines := scoredLines(r, c.in)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
	}
}
