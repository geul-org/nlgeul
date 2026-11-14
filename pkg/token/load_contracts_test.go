//ff:func feature=token type=reader control=sequence
//ff:what TestLoadContracts — the fused forms are listed set by set in the order of their lines; a line that is wrong ends the candidate with its error
package token

import "testing"
import "github.com/geul-org/nlgeul/data"

func TestLoadContracts(t *testing.T) {
	got, err := loadContracts([]data.Row{
		rowFor("준말", "이게", "이것", "(이)", "주격:이", "-", "-"),
		rowFor("오표기", "꺼", "거", "-", "보조사:는", "-", "꺼"),
		rowFor("준말", "그게", "그것", "(이)", "주격:이", "-", "-"),
	})
	if err != nil || len(got) != 2 || len(got[setContract]) != 2 || got[setContract][1].form != "그게" || got[setTypo][0].wrote != "꺼" {
		t.Fatal(got, err)
	}
	_, err = loadContracts([]data.Row{rowFor("준말", "이게", "이것", "(이)", "이", "-", "-")})
	errHas(t, err, "5번째 칸")
}
