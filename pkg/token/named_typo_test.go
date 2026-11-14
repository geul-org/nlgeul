//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNamedTypo — the entry found by a name is the right closing's, never the one that reads a closing written wrong under the same name
package token

import "testing"

func TestNamedTypo(t *testing.T) {
	ix := scanOf("x").ix
	right := ix.named('걸', "맺음:ㄹ걸")
	if right.wrote != "" || right.g == nil {
		t.Fatal(right)
	}
	found := false
	for _, e := range ix.byKey['걸'] {
		found = found || (e.name == "맺음:ㄹ걸" && e.wrote != "")
	}
	if found {
		t.Fatal("the entry written wrong is under the key of the right form's last letter, not 걸")
	}
	ix.byKey['x'] = append(ix.byKey['x'], &entry{name: "맺음:ㄹ걸", wrote: "ㄹ껄", form: "x"})
	if e := ix.named('x', "맺음:ㄹ걸"); e.wrote != "" || e.g != nil {
		t.Fatal("a name found only on an entry written wrong gives an entry with no graph", e)
	}
}
