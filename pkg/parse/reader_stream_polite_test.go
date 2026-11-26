//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamPolite — with a word table the old polite slots are read after a predicate's stem right before a closing (가온: 가+오+(ㄴ), 하옵니다: 하+옵+니다, 먹사오며, 받자옵고, 하시옵고, 먹으오며, 학교이옵니다), next to the other candidates; 하오니 keeps the tables' closing 오니 alone; with no word table none is read
package parse

import (
	"slices"
	"strings"
	"testing"
)

func TestReaderStreamPolite(t *testing.T) {
	r := streamReader(t, "words-hao.tsv")
	for _, c := range []struct{ sentence, want string }{
		{"가온 결과", "가[열쇠:동사]+오[선어말:오]+(ㄴ)[꾸밈:ㄴ]"},
		{"가올 결과", "가[열쇠:동사]+오[선어말:오]+(ㄹ)[꾸밈:ㄹ]"},
		{"그리 하옵니다.", "하[열쇠:동사]+옵[선어말:옵]+니다[맺음:니다]"},
		{"물을 먹사오며", "먹[열쇠:동사]+사오[선어말:사오]+며[맺음:며]"},
		{"물을 받자옵고", "받[열쇠:동사]+자옵[선어말:자옵]+고[맺음:고]"},
		{"그리 하시옵고", "하[열쇠:동사]+시[선어말:시]+옵[선어말:옵]+고[맺음:고]"},
		{"물을 먹으오며", "먹[열쇠:동사]+으[이음:으]+오[선어말:오]+며[맺음:며]"},
		{"여기는 학교이옵니다.", "학교[열쇠:체언]+이[서술격:이]+옵[선어말:옵]+니다[맺음:니다]"},
		{"그리 하오니", "하[열쇠:동사]+오니[맺음:오니]"},
	} {
		if lines := streamLines(r, c.sentence); !slices.Contains(lines, c.want) {
			t.Errorf("%s: no %q in %q", c.sentence, c.want, lines)
		}
	}
	if got := strings.Join(streamLines(r, "그리 하오니"), "\n"); strings.Contains(got, "선어말:오") {
		t.Fatal("오니 is the tables' closing", got)
	}
	if got := strings.Join(streamLines(streamReader(t, ""), "그리 하옵니다."), "\n"); strings.Contains(got, "선어말:옵") {
		t.Fatal("no word table, no polite slot", got)
	}
}
