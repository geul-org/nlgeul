//ff:func feature=parse type=reader control=sequence
//ff:what TestPoliteTakersOf — the polite slots that stand after a vowel take the link vowel 으 (오, 옵), the others do not; the table's own are 선어말:오 and 선어말:옵
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPoliteTakersOf(t *testing.T) {
	rows := []entry{{name: "선어말:오", after: rule.AfterVowel}, {name: "선어말:사오", after: rule.AfterBatchim}, {name: "선어말:옵", after: rule.AfterVowel}}
	if got := politeTakersOf(rows); !slices.Equal(got, []string{"선어말:오", "선어말:옵"}) {
		t.Fatal(got)
	}
	if politeTakersOf(nil) != nil || !slices.Equal(politeTakers, []string{"선어말:오", "선어말:옵"}) {
		t.Fatal(politeTakers)
	}
}
