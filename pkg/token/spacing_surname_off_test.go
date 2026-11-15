//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSpacingSurnameOff — the family name is read only with article 48 on: with no rule on, or with every rule on but 48, no candidate of the stream or of the old reading of 김군은, 오과장이다 or 남궁씨가 is named 성; with 48 alone on both readings have it
package token

import (
	"slices"
	"testing"
)

func TestSpacingSurnameOff(t *testing.T) {
	none := spacingReader(t, "words-surname.tsv")
	but := spacingReader(t, "words-surname.tsv", slices.DeleteFunc(slices.Clone(SpacingArticles), func(a string) bool { return a == "48" })...)
	only := spacingReader(t, "words-surname.tsv", "48")
	for _, s := range []string{"김군은 왔다", "오과장이다", "남궁씨가 왔다"} {
		if netHas(none, s, "[성:") || unitHas(none, s, "[성:") || netHas(but, s, "[성:") || unitHas(but, s, "[성:") {
			t.Error("the rule is off", s)
		}
		if !netHas(only, s, "[성:") || !unitHas(only, s, "[성:") {
			t.Error("the rule is on", s)
		}
	}
}
