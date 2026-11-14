//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNamesForm — a tail whose last piece is a modifying closing, or with a naming closing anywhere (particles after it too), makes a modifying or naming form; a modifying closing that is not the last piece, a tail of slots and a final closing, and no pieces at all do not
package token

import "testing"

func TestNamesForm(t *testing.T) {
	cases := []struct {
		pieces []piece
		want   bool
	}{
		{[]piece{{text: "은", names: "꾸밈:은"}}, true},
		{[]piece{{text: "었", names: "선어말:었"}, {text: "던", names: "꾸밈:던"}}, true},
		{[]piece{{text: "기", names: "이름:기"}, {text: "는", names: "보조사:는"}}, true},
		{[]piece{{text: "시", names: "선어말:시"}, {text: "ㅁ", names: "이름:ㅁ"}, {text: "을", names: "목적격:을"}}, true},
		{[]piece{{text: "ㄴ", names: "꾸밈:ㄴ"}, {text: "도", names: "보조사:도"}}, false},
		{[]piece{{text: "었", names: "선어말:었"}, {text: "다", names: "맺음:다"}}, false},
		{nil, false},
	}
	for i, c := range cases {
		if got := namesForm(c.pieces); got != c.want {
			t.Fatal(i, c.pieces, got)
		}
	}
}
