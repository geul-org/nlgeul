//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestMarkerOf — four columns from the column given are read into a closing: its form and lead, what must stand before it, its kinds and its ways of attaching; a bad form, a bad condition on what stands before, a lead form with such a condition, no kind and a name outside the sets are errors that name the column
package parse

import "testing"

func TestMarkerOf(t *testing.T) {
	m, err := MarkerOf(rowFor("앞", "면", "받침 뒤", "종결|연결", "으 붙음|요 안 됨|서술격 안 됨"), 1)
	want := Marker{Form: "면", After: afterBatchim, Kind: kindFinal | kindLink, Eu: true, NoYo: true, NoIda: true}
	if err != nil || m != want {
		t.Fatal(m, err)
	}
	m, err = MarkerOf(rowFor("ㄴ다", "-", "꾸밈", "서술격만|서술격 품음"), 0)
	want = Marker{Form: "다", Lead: jongN, Kind: kindAdn, NeedIda: true, SelfIda: true}
	if err != nil || m != want {
		t.Fatal(m, err)
	}
	m, err = MarkerOf(rowFor("어/아서", "-", "이름", "-"), 0)
	if err != nil || m.Lead != leadA || m.Form != "서" || m.Kind != kindNoun {
		t.Fatal(m, err)
	}
	cases := []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "-", "종결", "-"}, []string{"1번째 칸(꼴)", "한글 음절"}},
		{[]string{"다", "앞", "종결", "-"}, []string{"2번째 칸(앞말)", "쓸 수 있는 이름"}},
		{[]string{"ㄴ다", "받침 뒤", "종결", "-"}, []string{"2번째 칸(앞말)", "-를 적는다"}},
		{[]string{"다", "-", "-", "-"}, []string{"3번째 칸(갈래)", "하나는 있어야"}},
		{[]string{"다", "-", "동사", "-"}, []string{"3번째 칸(갈래)", "꾸밈, 연결, 이름, 종결"}},
		{[]string{"다", "-", "종결", "빠름"}, []string{"4번째 칸(조건)", "서술격 안 됨"}},
	}
	for _, c := range cases {
		_, err := MarkerOf(rowFor(c.row...), 0)
		errHas(t, err, c.parts...)
	}
}
