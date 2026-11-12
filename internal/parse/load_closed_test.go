//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestLoadClosed — the closed sets are listed in the order of their lines and the dependent nouns marked 붙여 읽음 are listed again under that mark; a set outside the four, a mark on a line of another set and a mark other than 붙여 읽음 are errors that name the column
package parse

import (
	"github.com/geul-org/nlgeul/data"
	"strings"
	"testing"
)

func TestLoadClosed(t *testing.T) {
	got, err := loadClosed([]data.Row{
		rowFor("의존명사", "것", "붙여 읽음"), rowFor("의존명사", "수", "-"), rowFor("의존명사", "줄", "붙여 읽음"),
		rowFor("보조 용언", "주", "-"), rowFor("관형사", "이", "-"), rowFor("접속 부사", "그리고", "-"),
	})
	if err != nil || strings.Join(got[boundSet], ",") != "것,수,줄" || strings.Join(got[joinedMark], ",") != "것,줄" ||
		strings.Join(got[helperSet], ",") != "주" || strings.Join(got[determinerSet], ",") != "이" || strings.Join(got[conjunctionSet], ",") != "그리고" {
		t.Fatal(got, err)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"동사", "가", "-"}, []string{"1번째 칸(꼴)", "쓸 수 있는 이름"}},
		{[]string{"관형사", "이", "붙여 읽음"}, []string{"3번째 칸(갈래)", "의존명사의 줄에만"}},
		{[]string{"의존명사", "것", "붙여씀"}, []string{"3번째 칸(갈래)", "붙여 읽음"}},
	} {
		_, err := loadClosed([]data.Row{rowFor(c.row...)})
		errHas(t, err, c.parts...)
	}
}
