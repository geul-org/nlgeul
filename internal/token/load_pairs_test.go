//ff:func feature=token type=reader control=sequence
//ff:what TestLoadPairs — the signs of a pair are read sign by sign; a sign written on two lines is an error that names the column, and so is a line that is wrong
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/data"
)

func TestLoadPairs(t *testing.T) {
	got, err := loadPairs([]data.Row{rowFor("(", "(", "열기", "-"), rowFor(")", "(", "닫기", "-")})
	if err != nil || len(got) != 2 || !got['('].opens || !got[')'].shuts || got[')'].kind != '(' {
		t.Fatal(got, err)
	}
	_, err = loadPairs([]data.Row{rowFor("(", "(", "열기", "-"), rowFor("(", "(", "열기", "-")})
	errHas(t, err, "1번째 칸(꼴)", "앞 줄에 이미 적혀 있다")
	_, err = loadPairs([]data.Row{rowFor("(", "(", "-", "-")})
	errHas(t, err, "3번째 칸", "열기와 닫기")
}
