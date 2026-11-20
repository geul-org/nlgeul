//ff:func feature=parse type=reader control=sequence
//ff:what TestNamingRows — of the rows given only those whose kinds take in the naming one are kept, in their order (one of two kinds as well); none given or none so give none; the gated closings give the naming closing 기 only
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNamingRows(t *testing.T) {
	rows := []Marker{{Form: "기", Kind: rule.KindNoun}, {Form: "다", Kind: rule.KindFinal}, {Form: "음", Kind: rule.KindNoun | rule.KindFinal}}
	if got := namingRows(rows); len(got) != 2 || got[0].Form != "기" || got[1].Form != "음" {
		t.Fatal(got)
	}
	if namingRows(nil) != nil || namingRows(rows[1:2]) != nil {
		t.Fatal("no naming rows")
	}
	if got := namingRows(gatedClosings); len(got) != 1 || got[0] != (Marker{Form: "기", Kind: rule.KindNoun, NoYo: true}) || !slices.Equal(nounHeads, []string{"이름:기"}) {
		t.Fatal(got, nounHeads)
	}
}
