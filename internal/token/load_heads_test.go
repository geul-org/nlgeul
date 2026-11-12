//ff:func feature=token type=reader control=sequence
//ff:what TestLoadHeads — the signs at the front of a number are named with the kind and the sign; a sign of more than one letter and a sign written on two lines are errors that name the column
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/data"
)

func TestLoadHeads(t *testing.T) {
	got, err := loadHeads([]data.Row{rowFor("-", "부호"), rowFor("$", "돈")})
	if err != nil || len(got) != 2 || got['-'] != "부호:-" || got['$'] != "돈:$" {
		t.Fatal(got, err)
	}
	_, err = loadHeads([]data.Row{rowFor("--", "부호")})
	errHas(t, err, "1번째 칸", "글자 하나가 아니다")
	_, err = loadHeads([]data.Row{rowFor("-", "부호"), rowFor("-", "돈")})
	errHas(t, err, "1번째 칸(꼴)", "앞 줄에 이미 적혀 있다")
}
