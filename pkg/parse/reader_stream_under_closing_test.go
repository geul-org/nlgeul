//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamUnderClosing — the stream reads 는 put under a closing's last syllable as ㄴ: 비해선 = 비하+(여)서+(는), 닫을진 = 닫+을지+(는), 가다간 = 가+다가+(는); 닫진 stays the shortened closing 진 alone
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamUnderClosing(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-under.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in, has, not string
	}{
		{"그것에 비해선 좋다.", "\n비하[열쇠:동사]+(여)서[맺음:어서(합침)]+(는)[보조사:는]\n", "비해선+"},
		{"어디에 닫을진 모른다.", "\n닫[열쇠:동사]+을지[맺음:을지]+(는)[보조사:는]\n", "닫을진+"},
		{"그렇게 가다간 늦는다.", "\n가[열쇠:동사]+다가[맺음:다가]+(는)[보조사:는]\n", "가다간+"},
		{"거기 닫진 않는다.", "\n닫[열쇠:동사]+진[맺음:진(지는)]\n", "지[맺음:지]+(는)"},
	} {
		got := streamText(r, c.in)
		if !strings.Contains(got, c.has) || strings.Contains(got, c.not) {
			t.Errorf("%q: want %q and no %q\n%s", c.in, c.has, c.not, got)
		}
	}
}
