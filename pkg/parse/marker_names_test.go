//ff:func feature=parse type=reader control=sequence
//ff:what TestMarkerNames — the names of the entries under every key of the index, each once; an index with none has no name
package parse

import "testing"

func TestMarkerNames(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{'소': {{name: "맺음:소"}, {name: "보조사:소"}}, '다': {{name: "맺음:다"}, {name: "맺음:소"}}}}
	got := ix.markerNames()
	if len(got) != 3 || !got["맺음:소"] || !got["보조사:소"] || !got["맺음:다"] {
		t.Fatal(got)
	}
	if got := (&index{}).markerNames(); len(got) != 0 {
		t.Fatal(got)
	}
}
