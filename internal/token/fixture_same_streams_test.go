//ff:func feature=token type=helper control=iteration dimension=1
//ff:what sameStreams — two streams of the same sentence give the same stretches to the end: the same place, number of ways, nodes and lambda
package token

import "testing"

func sameStreams(t *testing.T, text string, on, off *Stream) {
	t.Helper()
	for a, ok := on.Next(); ok; a, ok = on.Next() {
		b, _ := off.Next()
		if a.Region != b.Region || a.Ways.Cmp(b.Ways) != 0 || segmentLine(a) != segmentLine(b) || a.Lambda != b.Lambda {
			t.Fatal(text, segmentLine(a), segmentLine(b))
		}
	}
	if _, ok := off.Next(); ok {
		t.Fatal(text, "more stretches without the window")
	}
}
