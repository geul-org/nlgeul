//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestContractOf — a line of the fused forms' table is read into the form, the word, the particle as printed and as named (keyed by the last letter of the name), whether it is read only after a modifying closing and the letters it is wrongly written with; a set outside the names, a particle name that is not an index name and a place other than 꾸밈 뒤 are errors that name the column
package token

import "testing"

func TestContractOf(t *testing.T) {
	c, err := contractOf(rowFor("준말", "이게", "이것", "(이)", "주격:이", "-", "-"))
	if err != nil || c != (contract{"이게", "이것", "(이)", "주격:이", '이', false, ""}) {
		t.Fatal(c, err)
	}
	c, err = contractOf(rowFor("오표기", "꺼", "거", "-", "보조사:는", "꾸밈 뒤", "꺼"))
	if err != nil || c != (contract{"꺼", "거", "-", "보조사:는", '는', true, "꺼"}) {
		t.Fatal(c, err)
	}
	for _, r := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"그 밖", "이게", "이것", "(이)", "주격:이", "-", "-"}, []string{"1번째 칸", "준말", "오표기"}},
		{[]string{"준말", "이게", "이것", "(이)", "이", "-", "-"}, []string{"5번째 칸", "색인의 이름"}},
		{[]string{"준말", "이게", "이것", "(이)", "주격:", "-", "-"}, []string{"5번째 칸", "색인의 이름"}},
		{[]string{"준말", "이게", "이것", "(이)", "주격:이", "앞 뒤", "-"}, []string{"6번째 칸", "꾸밈 뒤"}},
	} {
		_, err := contractOf(rowFor(r.row...))
		errHas(t, err, r.parts...)
	}
}
