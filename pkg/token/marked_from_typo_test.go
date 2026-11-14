//ff:func feature=token type=reader control=sequence
//ff:what TestMarkedFromTypo — a tail with a marker written wrong does not count as letters read as markers of the index alone (the letters are a word of their own as well); the same tail written right does
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkedFromTypo(t *testing.T) {
	right := piece{text: "음", names: "이름:음", tail: rule.TailPred}
	wrong := piece{text: "음", names: "이름:음", tail: rule.TailPred, wrote: "슴"}
	got := markedFrom([]tail{{start: 1, pieces: []piece{right}}, {start: 2, pieces: []piece{wrong}}, {start: 3, pieces: []piece{right, wrong}}})
	if len(got) != 1 || !got[1] {
		t.Fatal(got)
	}
}
