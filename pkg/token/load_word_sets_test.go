//ff:func feature=token type=reader control=sequence
//ff:what TestLoadWordSets — the sets given are listed in the order of their lines and a set given that no line names has an empty list; the lines of a set pkg/parse reads are left out; a set of neither is an error that names the column and the names that may be written
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func TestLoadWordSets(t *testing.T) {
	got, err := loadWordSets([]data.Row{
		rowFor("대명사", "네"), rowFor("보조 용언", "주"), rowFor("있다·없다·계시다", "있"),
		rowFor("대명사", "내"), rowFor("관형사", "이"),
	}, setPronounWord, setExistStem, setAskWord)
	if err != nil || len(got) != 3 || strings.Join(got[setPronounWord], ",") != "네,내" || strings.Join(got[setExistStem], ",") != "있" ||
		got[setAskWord] != nil {
		t.Fatal(got, err)
	}
	if _, ok := got[auxSet]; ok {
		t.Fatal("a set pkg/parse reads", got)
	}
	_, err = loadWordSets([]data.Row{rowFor("동사", "가")}, setPronounWord)
	errHas(t, err, "1번째 칸(꼴)", "쓸 수 있는 이름", "대명사", "보조 용언")
}
