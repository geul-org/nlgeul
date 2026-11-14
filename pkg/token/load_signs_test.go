//ff:func feature=token type=reader control=sequence
//ff:what TestLoadSigns — the candidates of a sign are listed in the order of their lines; a sign of more than one letter is an error
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/data"
)

func TestLoadSigns(t *testing.T) {
	got, err := loadSigns([]data.Row{rowFor(",", "쉼표"), rowFor("?", "물음표"), rowFor(",", "반점")})
	if err != nil || len(got) != 2 || strings.Join(got[','], ",") != "쉼표,반점" || got['?'][0] != "물음표" {
		t.Fatal(got, err)
	}
	_, err = loadSigns([]data.Row{rowFor(",,", "쉼표")})
	errHas(t, err, "1번째 칸(꼴)", "글자 하나가 아니다")
}
