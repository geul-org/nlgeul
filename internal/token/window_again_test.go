//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestWindowAgain — a stretch a window cut off is read again as it was first read: the first ways of the stretch (up to five) are the same words in the same order with the same lambdas, for the stretches of made-up long units, the first of a unit and the ones after it
package token

import "testing"

func TestWindowAgain(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, text := range []string{redoText, "(고양이)" + redoText, "아버지/가방에들어가신다·아버지가방에들어가신다아버지가방에들어가신다아버지가방에들어가신다"} {
		st := r.Stream(text)
		for st.fill(); len(st.queue) > 0; st.fill() {
			checkRedo(t, st.queue[0])
			st.Next()
		}
		if st.stats.cuts == 0 {
			t.Fatal(text, "no cut")
		}
	}
}
