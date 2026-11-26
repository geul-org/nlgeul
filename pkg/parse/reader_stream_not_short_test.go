//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNotShort — with a word table the adjective 아니 takes the shortened quotations the copula takes (아니란: 아니+란, 아니래, 아니랄까), as it takes 라는; a verb's stem keeps them, and another adjective does not take them
package parse

import (
	"slices"
	"strings"
	"testing"
)

func TestReaderStreamNotShort(t *testing.T) {
	r := streamReader(t, "words-hao.tsv")
	for _, c := range []struct{ sentence, want string }{
		{"학교가 아니란 결과", "아니[열쇠:형용사]+란[꾸밈:란(라는)]"},
		{"학교가 아니래.", "아니[열쇠:형용사]+래[맺음:래(라고 해)]"},
		{"학교가 아니랄까.", "아니[열쇠:형용사]+랄까[맺음:랄까(라고 할까)]"},
		{"학교에 가란 결과", "가[열쇠:동사]+란[꾸밈:란(라는)]"},
	} {
		if lines := streamLines(r, c.sentence); !slices.Contains(lines, c.want) {
			t.Errorf("%s: no %q in %q", c.sentence, c.want, lines)
		}
	}
	if lines := streamLines(r, "작란 결과"); strings.Contains(strings.Join(lines, "\n"), "작[열쇠:형용사]+란") {
		t.Fatal("another adjective", lines)
	}
}
