//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestEndingOf — a line of the closing table is read into the closing and its count; the flags are kept as bits; an error of the closing's columns or a count that is not a number is passed on with its column
package parse

import "testing"

func TestEndingOf(t *testing.T) {
	e, err := endingOf(rowFor("ㄴ다", "-", "종결|연결", "으 붙음|요 안 됨", "12"))
	want := ending{"다", jongN, afterAny, kindFinal | kindLink, fEu | fNoYo, 12}
	if err != nil || e != want {
		t.Fatal(e, err)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "-", "종결", "-", "1"}, []string{"1번째 칸(꼴)"}},
		{[]string{"다", "-", "종결", "-", "많이"}, []string{"5번째 칸(횟수)", "0 이상의 수가 아니다"}},
	} {
		_, err := endingOf(rowFor(c.row...))
		errHas(t, err, c.parts...)
	}
}
