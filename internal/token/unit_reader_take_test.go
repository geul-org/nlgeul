//ff:func feature=token type=helper control=sequence
//ff:what TestUnitReaderTake — the trace lines made since they were last taken come once, in the order made, and what is made after is another taking
package token

import "testing"

func TestUnitReaderTake(t *testing.T) {
	rd := &unitReader{notes: []string{"a", "b"}}
	got := rd.take()
	rd.notes = append(rd.notes, "c")
	if len(got) != 2 || got[0] != "a" || len(rd.take()) != 1 || rd.take() != nil {
		t.Fatal(got)
	}
}
