//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestAfterCopula — a tail whose second piece is an auxiliary of copulaBars (하 of 어하다) gives that auxiliary's name, any other tail what firstName gives: its closing, a slot, or the first piece's name; a tail of one piece, or whose second piece is another auxiliary (지), gives the first piece's
package token

import "testing"

func TestAfterCopula(t *testing.T) {
	eo := piece{text: "어", names: "맺음:어(합침)"}
	for _, c := range []struct {
		t    tail
		want string
	}{
		{tail{pieces: []piece{eo, {text: "하", names: "보조:하"}, {text: "는", names: "꾸밈:는"}}}, "보조:하"},
		{tail{pieces: []piece{{text: "(어)", names: "맺음:어(합침)"}, {text: "하", names: "보조:하"}}}, "보조:하"},
		{tail{pieces: []piece{eo, {text: "지", names: "보조:지"}, {text: "는", names: "꾸밈:는"}}}, "맺음:어(합침)"},
		{tail{pieces: []piece{eo}}, "맺음:어(합침)"},
		{tail{pieces: []piece{{text: "하", names: "보조:하"}}}, "보조:하"},
		{tail{pieces: []piece{{text: "었", names: "선어말:었"}, {text: "다", names: "맺음:다"}}, close: &entry{name: "맺음:다"}}, "선어말:었"},
		{tail{pieces: []piece{{text: "다", names: "맺음:다"}}, close: &entry{name: "맺음:다"}}, "맺음:다"},
		{tail{pieces: []piece{eo, {text: "서", names: "맺음:서"}}, close: &entry{name: "맺음:어서"}}, "맺음:어(합침)"},
	} {
		if got := afterCopula(c.t); got != c.want {
			t.Errorf("%+v: %s, want %s", c.t.pieces, got, c.want)
		}
	}
}
