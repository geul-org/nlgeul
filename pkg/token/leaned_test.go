//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestLeaned — the ways before a quotation mark that leans to opening put those with none of its kind open first, before one that leans to closing those with one open first, each part in the order it was; with no lean (a bracket, a mark between letters) the ways are as they were
package token

import (
	"reflect"
	"testing"
)

func TestLeaned(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text string
		at   int
		open []string
		want []string
	}{
		{"가 '봄", 2, []string{"'", "(", "", "('"}, []string{"(", "", "'", "('"}},
		{"간다' 가", 2, []string{"", "'", "(", "('"}, []string{"'", "('", "", "("}},
		{"가'나", 1, []string{"'", ""}, []string{"'", ""}},
		{"봄) 가", 1, []string{"", "("}, []string{"", "("}},
	} {
		cut := newCutter(ix, c.text, false)
		cut.open = c.open
		if got := cut.leaned(c.at, pairRows[cut.r[c.at]]); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%q at %d with %q: %q, want %q", c.text, c.at, c.open, got, c.want)
		}
	}
}
