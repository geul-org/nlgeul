//ff:func feature=token type=reader control=sequence
//ff:what TestMarkerRow — a line of this layer's closing table is read as the closing tables write one; an error of its columns is passed on
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerRow(t *testing.T) {
	m, err := markerRow(rowFor("되", "-", "연결", "-"))
	if err != nil || m != (parse.Marker{Form: "되", Kind: rule.KindLink}) {
		t.Fatal(m, err)
	}
	_, err = markerRow(rowFor("되", "-", "-", "-"))
	errHas(t, err, "표 t.tsv 2번째 줄 3번째 칸(갈래)")
}
