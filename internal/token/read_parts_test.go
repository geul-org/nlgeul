//ff:func feature=token type=helper control=sequence
//ff:what TestReadParts — a candidate is given once for each part, in the parts' order, with its word k holding that part
package token

import "testing"

func TestReadParts(t *testing.T) {
	br := Branch{Words: []Eojeol{{Parts: []Part{{Text: "x"}}}}}
	got := readParts(br, 0, []Part{{Text: "a"}, {Text: "b"}})
	if len(got) != 2 || got[0].Words[0].Parts[0].Text != "a" || got[1].Words[0].Parts[0].Text != "b" {
		t.Fatal(got)
	}
	if got = readParts(br, 0, nil); len(got) != 0 {
		t.Fatal("no part", got)
	}
}
