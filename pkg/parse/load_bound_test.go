//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestLoadBound — the words that bear the set's name are listed in the order of their lines and those marked 붙여 읽음 again under that mark; a word with no name is left out of both, whatever the columns loadBoundNouns reads hold; a name that is not the set's, a mark other than 붙여 읽음 and a mark on a line with no name are errors that name the column
package parse

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func TestLoadBound(t *testing.T) {
	got, err := loadBound([]data.Row{
		rowFor("것", "의존명사", "붙여 읽음", "꾸밈 뒤", "-"), rowFor("중", "-", "-", "체언 뒤", "-"),
		rowFor("간", "의존명사", "-", "체언 뒤", "-"), rowFor("줄", "의존명사", "붙여 읽음", "꾸밈 뒤", "-"),
	})
	if err != nil || len(got) != 2 || strings.Join(got[boundSet], ",") != "것,간,줄" || strings.Join(got[joinedMark], ",") != "것,줄" {
		t.Fatal(got, err)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"것", "관형사", "-", "꾸밈 뒤", "-"}, []string{"2번째 칸(앞말)", "의존명사 또는 -"}},
		{[]string{"것", "의존명사", "붙여씀", "꾸밈 뒤", "-"}, []string{"3번째 칸(갈래)", "붙여 읽음"}},
		{[]string{"중", "-", "붙여 읽음", "체언 뒤", "-"}, []string{"3번째 칸(갈래)", "이름표가 의존명사인 줄에만"}},
	} {
		_, err := loadBound([]data.Row{rowFor(c.row...)})
		errHas(t, err, c.parts...)
	}
}
