//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNumeralsLong — a run of Hangul numerals has no length limit: before 명 every run that ends there is read, up to the seven letters of 삼천오백육십칠, each dearer than the one inside it; the run is looked for no further back than the letters of numerals go (곰 stops it); end to end, with a word table and without, 삼천오백육십칠명 is read first as the number and its unit
package token

import (
	"strings"
	"testing"
)

func TestNumeralsLong(t *testing.T) {
	got := scanOf("삼천오백육십칠명").numerals(7)
	want := []string{"칠", "십칠", "육십칠", "백육십칠", "오백육십칠", "천오백육십칠", "삼천오백육십칠"}
	if len(got) != len(want) {
		t.Fatal(got)
	}
	for i, n := range got {
		if n.text != want[i] || (i > 0 && n.pref >= got[i-1].pref) {
			t.Errorf("%d: %+v, want %s", i, n, want[i])
		}
	}
	if short := scanOf("곰삼십명").numerals(3); len(short) != 2 || short[1].text != "삼십" {
		t.Fatal("a letter that is no numeral ends the run", short)
	}
	for _, opt := range []Options{{}, {Words: "testdata/words.tsv"}} {
		r, _ := New(opt)
		if lines := r.Lines("삼천오백육십칠명", false); len(lines) < 2 || !strings.HasPrefix(lines[1], "  삼천오백육십칠[수]+명[단위:명]\t") {
			t.Errorf("삼천오백육십칠명 (%q)\n%s", opt.Words, strings.Join(lines, "\n"))
		}
	}
}
