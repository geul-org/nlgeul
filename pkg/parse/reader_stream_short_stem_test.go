//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamShortStem — the stream reads the shortened forms of a fused word: 이걸로 = 이것+으로, and 데다 right after a modifying form = 데+(에)다
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamShortStem(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-under.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in, has, not string
	}{
		{"이걸로 씻는다.", "\n이것[열쇠:체언]+으로[부사격:으로]\n", "이것[열쇠:체언]+(으)로"},
		{"그걸로 씻는다.", "\n그것[열쇠:체언]+으로[부사격:으로]\n", "그것[열쇠:체언]+(으)로"},
		{"바람이 부는 데다 비가 온다.", "\n데[의존명사:데]+(에)다[부사격:에다]\n", "데[의존명사:데]+(에)다가"},
		{"바람이 부는 데다가 비가 온다.", "\n데[의존명사:데]+(에)다가[부사격:에다가]\n", "데[의존명사:데]+다가"},
	} {
		got := streamText(r, c.in)
		if !strings.Contains(got, c.has) || strings.Contains(got, c.not) {
			t.Errorf("%q: want %q and no %q\n%s", c.in, c.has, c.not, got)
		}
	}
}
