//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestFirstName — a tail that begins with its closing gives the closing's name (one of the names of its first piece too), so does one that begins with the closing written out as its parts (죠 as 지+(요), 세요 as 시+(어)+요, 잖 whose closing is 지); a slot or the link vowel before the closing gives its own name, and a tail with no closing its first piece's; end to end, with a word table and without, 학생이었습니다, 학생이었다 and 학생이었는데 are read first as 학생+이+었+… with 이 the copula's stem, and 의사였다 has a candidate with the copula's stem before the fused slot
package token

import (
	"strings"
	"testing"
)

func TestFirstName(t *testing.T) {
	da, joa := &entry{name: "맺음:다"}, &entry{name: "맺음:죠"}
	past := piece{text: "었", names: "선어말:었"}
	for _, c := range []struct {
		t    tail
		want string
	}{
		{tail{pieces: []piece{past, {text: "다", names: "맺음:다"}}}, "선어말:었"},
		{tail{pieces: []piece{{text: "다", names: "맺음:다"}}}, "맺음:다"},
		{tail{pieces: []piece{{text: "습니다", names: "맺음:습니다"}}, close: &entry{name: "맺음:습니다"}}, "맺음:습니다"},
		{tail{pieces: []piece{{text: "나", names: "맺음:나"}}, close: &entry{name: "맺음:나"}}, "맺음:나"},
		{tail{pieces: []piece{past, {text: "습니다", names: "맺음:습니다"}}, close: &entry{name: "맺음:습니다"}}, "선어말:었"},
		{tail{pieces: []piece{{text: "(었)", names: "선어말:ㅆ"}, {text: "다", names: "맺음:다"}}, base: '이', close: da}, "선어말:ㅆ"},
		{tail{pieces: []piece{{text: "으", names: euName}, {text: "며", names: "맺음:며"}}, close: &entry{name: "맺음:며"}}, euName},
		{tail{pieces: []piece{{text: "지", names: jiName}, {text: "(요)", names: yoName}}, close: joa}, "맺음:죠"},
		{tail{pieces: []piece{past, {text: "지", names: jiName}, {text: "(요)", names: yoName}}, close: joa}, "선어말:었"},
		{tail{pieces: []piece{{text: "시", names: honorName}, {text: "(어)", names: vowelName}, {text: "요", names: yoName}}, close: &entry{name: "맺음:세요"}}, "맺음:세요"},
		{tail{pieces: []piece{{text: "시", names: honorName}, {text: "다", names: "맺음:다"}}, close: da}, honorName},
		{tail{pieces: []piece{{text: "지", names: jiName}, {text: "(않)", names: "부정:않"}, {text: "아", names: "맺음:아"}}, close: &entry{name: jiName}}, jiName},
	} {
		if got := firstName(c.t); got != c.want {
			t.Errorf("%+v: %s, want %s", c.t.pieces, got, c.want)
		}
	}
	bare, _ := New(Options{})
	lexed, _ := New(Options{Words: "testdata/words.tsv"})
	for _, c := range []struct {
		r     *Reader
		in    string
		want  string
		first bool
	}{
		{bare, "학생이었습니다", "  학생+이[서술격:이]+었[선어말:었]+습니다[맺음:습니다]\t", true},
		{lexed, "학생이었습니다", "  학생+이[서술격:이]+었[선어말:었]+습니다[맺음:습니다]\t", true},
		{bare, "학생이었다", "  학생+이[서술격:이]+었[선어말:었]+다[맺음:다]\t", true},
		{lexed, "학생이었다", "  학생+이[서술격:이]+었[선어말:었]+다[맺음:다]\t", true},
		{bare, "학생이었는데", "  학생+이[서술격:이]+었[선어말:었]+는데[맺음:는데]\t", true},
		{lexed, "학생이었는데", "  학생+이[서술격:이]+었[선어말:었]+는데[맺음:는데]\t", true},
		{bare, "의사였다", "  의사+이[서술격:이]+(었)[선어말:ㅆ]+다[맺음:다]\t", false},
		{lexed, "의사였다", "  의사+이[서술격:이]+(었)[선어말:ㅆ]+다[맺음:다]\t", false},
	} {
		lines := c.r.Lines(c.in, false)
		if got := strings.Join(lines, "\n"); !strings.Contains(got, c.want) || (c.first && !strings.HasPrefix(lines[1], c.want)) {
			t.Errorf("%s (word table: %v): no %q (first: %v)\n%s", c.in, c.r == lexed, c.want, c.first, got)
		}
	}
}
