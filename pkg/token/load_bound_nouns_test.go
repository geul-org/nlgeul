//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestLoadBoundNouns — the words of each place and of each use are listed in the order of their lines, a word under each name written on its line; a place or a use that is not of its set, a name written twice and a line with no place are errors that name the column
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func TestLoadBoundNouns(t *testing.T) {
	got, err := loadBoundNouns([]data.Row{
		rowFor("거", "의존명사", "붙여 읽음", "꾸밈 뒤", "꾸밈 맺음 뒤 붙여 씀"), rowFor("수", "의존명사", "붙여 읽음", "꾸밈 뒤", "꾸밈 뒤가 아니면 체언"),
		rowFor("때", "의존명사", "붙여 읽음", "꾸밈 뒤|체언 뒤", "-"), rowFor("중", "-", "-", "체언 뒤", "-"),
		rowFor("만", "의존명사", "-", "꾸밈 뒤|시간 뒤", "-"), rowFor("바", "-", "-", "꾸밈 뒤", "-"),
	})
	if err != nil || len(got) != 5 || strings.Join(got[adnPlace], ",") != "거,수,때,만,바" || strings.Join(got[nominalPlace], ",") != "때,중" ||
		strings.Join(got[timePlace], ",") != "만" || strings.Join(got[useAdnOnto], ",") != "거" || strings.Join(got[useAdnBound], ",") != "수" {
		t.Fatal(got, err)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"바", "-", "-", "관형어 뒤", "-"}, []string{"4번째 칸(조건)", "쓸 수 있는 이름", "꾸밈 뒤"}},
		{[]string{"바", "-", "-", "꾸밈 뒤|꾸밈 뒤", "-"}, []string{"4번째 칸(조건)", "두 번"}},
		{[]string{"바", "-", "-", "-", "-"}, []string{"4번째 칸(조건)", "자리를 하나 이상"}},
		{[]string{"바", "-", "-", "꾸밈 뒤", "붙여 읽음"}, []string{"5번째 칸(횟수)", "쓸 수 있는 이름", "꾸밈 맺음 뒤 붙여 씀"}},
	} {
		_, err := loadBoundNouns([]data.Row{rowFor(c.row...)})
		errHas(t, err, c.parts...)
	}
}
