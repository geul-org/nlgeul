//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderLinesShortClosing — end to end with the small table of testdata: the shortened closings 달까, ㄴ달까, 는달까 and 랄까 are read after the stems they follow (좋달까, 간달까, 먹는달까, 학생이랄까), and with no word table no reading holds their names
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesShortClosing(t *testing.T) {
	with, err := New(Options{Words: "testdata/words-short.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	without, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ sentence, want, name string }{
		{"좋달까", "좋[열쇠:형용사]+달까[맺음:달까(다고 할까)]", "맺음:달까(다고 할까)"},
		{"간달까", "가[열쇠:동사]+(ㄴ)달까[맺음:ㄴ달까(ㄴ다고 할까)]", "맺음:ㄴ달까(ㄴ다고 할까)"},
		{"먹는달까", "먹[열쇠:동사]+는달까[맺음:는달까(는다고 할까)]", "맺음:는달까(는다고 할까)"},
		{"학생이랄까", "학생[열쇠:체언]+이[서술격:이]+랄까[맺음:랄까(라고 할까)]", "맺음:랄까(라고 할까)"},
	}
	for _, c := range cases {
		got := strings.Join(with.Lines(c.sentence, false), "\n")
		if !strings.Contains(got, "  "+c.want+"\t") {
			t.Errorf("%s: missing %q in\n%s", c.sentence, c.want, got)
		}
		if bare := strings.Join(without.Lines(c.sentence, false), "\n"); strings.Contains(bare, c.name) {
			t.Errorf("%s: %q read with no word table:\n%s", c.sentence, c.name, bare)
		}
	}
}
