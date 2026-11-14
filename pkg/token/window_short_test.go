//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestWindowShort — a unit that is short (or has no word every way agrees on before its end) is read the same with the window on and off: the stretches, the ways, and the numbers of ways; the sentences of the tests are read the same
package token

import (
	"testing"
)

func TestWindowShort(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, text := range append(netSentences(), ontoSentences()...) {
		on, off := r.Stream(text), r.Stream(text)
		off.ix = cloneNoWindow(r.ix)
		sameStreams(t, text, on, off)
	}
}
