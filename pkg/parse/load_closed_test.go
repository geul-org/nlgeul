//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestLoadClosed — the closed sets are listed in the order of their lines; the line of a set outside the four is left out; a line of the dependent nouns is an error that names the column and the table they are written in
package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
	"strings"
	"testing"
)

func TestLoadClosed(t *testing.T) {
	got, err := loadClosed([]data.Row{
		rowFor("보조 용언", "주"), rowFor("관형사", "이"), rowFor("보조 용언", "보"), rowFor("접속 부사", "그리고"),
		rowFor("대명사", "네"), rowFor("지시어", "이"), rowFor("지시어", "그"),
	})
	if err != nil || len(got) != 4 || strings.Join(got[pointerSet], ",") != "이,그" || strings.Join(got[helperSet], ",") != "주,보" || strings.Join(got[determinerSet], ",") != "이" ||
		strings.Join(got[conjunctionSet], ",") != "그리고" {
		t.Fatal(got, err)
	}
	for _, rows := range [][]data.Row{{rowFor("의존명사", "것")}, {rowFor("관형사", "이"), rowFor("의존명사", "것")}} {
		_, err := loadClosed(rows)
		errHas(t, err, "1번째 칸(꼴)", "의존명사는", "bound-nouns.tsv")
	}
}
