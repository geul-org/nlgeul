//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestLoadWordSets — the sets given are listed in the order of their lines and a set given that no line names has an empty list; the lines of a set pkg/parse reads are left out; a set of neither and a mark on a line of a set given are errors that name the column
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func TestLoadWordSets(t *testing.T) {
	got, err := loadWordSets([]data.Row{
		rowFor("대명사", "네", "-"), rowFor("의존명사", "것", "붙여 읽음"), rowFor("구어 꼴", "거", "-"),
		rowFor("대명사", "내", "-"), rowFor("관형사", "이", "-"),
	}, setPronounWord, setSpoken, setAskWord)
	if err != nil || len(got) != 3 || strings.Join(got[setPronounWord], ",") != "네,내" || strings.Join(got[setSpoken], ",") != "거" ||
		got[setAskWord] != nil {
		t.Fatal(got, err)
	}
	if _, ok := got[boundSet]; ok {
		t.Fatal("a set pkg/parse reads", got)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"동사", "가", "-"}, []string{"1번째 칸(꼴)", "쓸 수 있는 이름", "대명사", "의존명사"}},
		{[]string{"대명사", "네", "붙여 읽음"}, []string{"3번째 칸(갈래)", "-를 적는다"}},
	} {
		_, err := loadWordSets([]data.Row{rowFor(c.row...)}, setPronounWord)
		errHas(t, err, c.parts...)
	}
}
