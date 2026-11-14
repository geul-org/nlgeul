//ff:func feature=token type=helper control=sequence
//ff:what TestSpanBefore — a span comes before another by its start, and with the same start by its end
package token

import "testing"

func TestSpanBefore(t *testing.T) {
	if !spanBefore(Span{0, 5}, Span{1, 2}) || spanBefore(Span{1, 2}, Span{0, 5}) {
		t.Fatal("start decides")
	}
	if !spanBefore(Span{1, 2}, Span{1, 3}) || spanBefore(Span{1, 3}, Span{1, 3}) || spanBefore(Span{1, 3}, Span{1, 2}) {
		t.Fatal("end decides")
	}
}
