//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSplitsOf — the closings of splits are written out with a word table and without (죠, 십시오); 시오 and 으시오 only with one, as the honorific and the 하오체 오, and without one they stay whole; a closing of neither is written whole; end to end, without a word table 가시오 and 먹으시오 read 시오 and 으시오 whole and nothing names 맺음:오, the graph that index has not, while with a word table they read 시+오
package token

import (
	"strings"
	"testing"
)

func TestSplitsOf(t *testing.T) {
	bare, lexed := &index{}, &index{lexed: true}
	for _, name := range []string{"맺음:죠", "맺음:십시오"} {
		if len(bare.splitsOf(name)) == 0 || len(lexed.splitsOf(name)) == 0 {
			t.Errorf("%s is written out with a word table and without", name)
		}
	}
	for name, n := range map[string]int{"맺음:시오": 2, "맺음:으시오": 3} {
		parts := lexed.splitsOf(name)
		if bare.splitsOf(name) != nil || len(parts) != n || parts[n-2].name != honorName || parts[n-1].name != "맺음:오" {
			t.Errorf("%s: %v without a word table, %v with one", name, bare.splitsOf(name), parts)
		}
	}
	if bare.splitsOf("맺음:다") != nil || lexed.splitsOf("맺음:다") != nil {
		t.Error("a closing of neither table is written whole")
	}

	plain, _ := New(Options{})
	words, _ := New(Options{Words: "testdata/words-hao.tsv"})
	for _, c := range []struct {
		r        *Reader
		in, want string
		gone     string
	}{
		{plain, "어서 가시오.", "  가+시오[맺음:시오]\t", "맺음:오]"},
		{plain, "먹으시오", "  먹+으시오[맺음:으시오]\t", "맺음:오]"},
		{words, "어서 가시오.", "  가[열쇠:동사]+시[선어말:시]+오[맺음:오]\t", "맺음:시오]"},
		{words, "먹으시오", "  먹[열쇠:동사]+으[이음:으]+시[선어말:시]+오[맺음:오]\t", "맺음:으시오]"},
	} {
		got := strings.Join(c.r.Lines(c.in, false), "\n")
		if !strings.Contains(got, c.want) || strings.Contains(got, c.gone) {
			t.Errorf("%s (word table: %v): want %q, no %q\n%s", c.in, c.r == words, c.want, c.gone, got)
		}
	}
}
