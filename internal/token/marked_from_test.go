//ff:func feature=token type=reader control=sequence
//ff:what TestMarkedFrom — where a tail of markers starts is kept; not for a tail fused into the stem, one that names its stem, one after a number, one that begins with a piece that does not show, or a sign
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkedFrom(t *testing.T) {
	p := piece{text: "에서", tail: rule.TailCase}
	got := markedFrom([]tail{
		{start: 1, pieces: []piece{p}},
		{start: 2, pieces: []piece{p}, base: '가'},
		{start: 3, pieces: []piece{p}, only: "아니"},
		{start: 4, pieces: []piece{p}, as: numName},
		{start: 5, pieces: []piece{{text: "(이)", unseen: true, tail: rule.TailCopula}}},
		{start: 6, pieces: []piece{{text: "."}}},
		{start: 7},
	})
	if len(got) != 1 || !got[1] {
		t.Fatal(got)
	}
}
