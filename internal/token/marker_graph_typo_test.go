//ff:func feature=token type=rule control=sequence
//ff:what TestMarkerGraphTypo — the graph of an entry that reads a closing written wrong has one edge more than the same entry's (the rebuttal of a misspelt form), and the edge is the same whatever the other exceptions are
package token

import "testing"

func TestMarkerGraphTypo(t *testing.T) {
	right := entry{name: "맺음:고", form: "고", class: classClosing}
	wrong := right
	wrong.wrote, wrong.right = "고", "고"
	if a, b := markerGraph(&right, []string{"그리고"}, false), markerGraph(&wrong, []string{"그리고"}, false); b != a+1 || right.g == nil || wrong.g == nil {
		t.Fatal(a, b)
	}
	if a, b := markerGraph(&right, nil, false), markerGraph(&wrong, nil, false); b != a+1 {
		t.Fatal(a, b)
	}
}
