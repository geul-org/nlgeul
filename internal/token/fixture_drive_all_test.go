//ff:func feature=token type=helper control=iteration dimension=3
//ff:what driveAll — reads by the net, with every word table of the tests, the sentences of the tests (netSentences, ontoSentences) and a few made-up units that hold a long known word, words bound by a slash or a hyphen, a wrong way to write a long word, a span of code and a sign of a pair with letters after it; each gives a stretch or more
package token

import (
	"strings"
	"testing"
)

func driveAll(t *testing.T) {
	t.Helper()
	units := append(netSentences(), ontoSentences()...)
	units = append(units,
		"대한민국임시정부수립기념일을마신다", "고양이가대한민국임시정부수립기념일을마신다",
		"대한민국임시정부/수립기념일을마신다", "대한민국임시정부-수립기념일을마신다",
		"대한민국임시정부수립기렴일을마신다", "대한민국임시정부수립했다", "대한민국임시정부를잘들었다",
		"가나다라마바사아자차카를마신다", "가나다라마바사아자차카타파하거너더러를", "abc/def-ghi를마신다",
		"`"+strings.Repeat("가", 30)+"`를마신다", ")"+strings.Repeat("고양이가마당에서물을마신다", 2),
		"(고양이)대한민국임시정부수립기념일을마신다", "물/마당-고양이가마신다")
	readers := netReaders(t)
	readers["far"] = streamReader(t, "words-far.tsv")
	for name, r := range readers {
		for _, s := range units {
			if len(netRegions(r, s)) == 0 {
				t.Fatal(name, s, "no stretch")
			}
		}
	}
}
