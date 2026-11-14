//ff:func feature=token type=reader control=sequence
//ff:what TestStreamNotes — Notes gives the lines kept since it was last called, and then nothing
package token

import "testing"

func TestStreamNotes(t *testing.T) {
	st := &Stream{notes: []string{"a", "b"}}
	if got := st.Notes(); len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatal(got)
	}
	if got := st.Notes(); got != nil {
		t.Fatal(got)
	}
}
