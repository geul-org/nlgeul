//ff:func feature=token type=reader control=sequence
//ff:what TestPref — a counted marker's preference is the log of its share; an uncounted one costs its syllables as stem syllables
package token

import (
	"math"
	"testing"
)

func TestPref(t *testing.T) {
	ix := &index{total: 1000}
	if math.Abs(ix.pref(10, "은")+2) > 1e-9 || math.Abs(ix.pref(0, "지만")-2*math.Log10(1.0/syllables)) > 1e-9 || ix.pref(0, "") != 0 {
		t.Fatal(ix.pref(10, "은"), ix.pref(0, "지만"))
	}
}
