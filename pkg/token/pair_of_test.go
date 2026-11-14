//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestPairOf — a line of the pairs' table is read into the sign and its row: the opening sign of its kind, whether it opens and closes, whether it is unsure and what keeps it from being a sign of a pair; a sign or kind of more than one letter, no role and a name outside the sets are errors that name the column
package token

import "testing"

func TestPairOf(t *testing.T) {
	sign, p, err := pairOf(rowFor(")", "(", "닫기", "-"))
	if err != nil || sign != ')' || p != (pairRow{kind: '(', shuts: true}) {
		t.Fatal(sign, p, err)
	}
	sign, p, err = pairOf(rowFor("'", "'", "열기|닫기", "헷갈림|낱말 속 안 됨|숫자 뒤 안 됨"))
	want := pairRow{kind: '\'', opens: true, shuts: true, unsure: true, bars: barWord | barDigit}
	if err != nil || sign != '\'' || p != want {
		t.Fatal(sign, p, err)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"((", "(", "열기", "-"}, []string{"1번째 칸", "글자 하나가 아니다"}},
		{[]string{"(", "((", "열기", "-"}, []string{"2번째 칸", "글자 하나가 아니다"}},
		{[]string{"(", "(", "-", "-"}, []string{"3번째 칸", "열기와 닫기"}},
		{[]string{"(", "(", "열림", "-"}, []string{"3번째 칸", "닫기, 열기"}},
		{[]string{"(", "(", "열기", "그냥"}, []string{"4번째 칸", "띄어 쓴 자리 안 됨"}},
	} {
		_, _, err := pairOf(rowFor(c.row...))
		errHas(t, err, c.parts...)
	}
}
