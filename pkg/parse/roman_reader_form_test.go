//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestRomanReaderForm — a Roman numeral letter of Unicode standing alone has two candidates, the number first and the letter's shape after it (Ⅲ[수:로마십진], Ⅲ[로마 숫자:수]); letters in a row and a letter with a unit or a particle on it are the number alone (ⅩⅢ, Ⅷ세, Ⅲ의); a full-width letter or digit alone keeps its reading as a run and gets the shape after it (Ａ[전각:로마자], １[전각:수])
package parse

import (
	"strings"
	"testing"
)

func TestRomanReaderForm(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for text, want := range map[string][]string{
		"Ⅲ": {"Ⅲ[수:로마십진]", "Ⅲ[로마 숫자:수]"},
		"ⅳ": {"ⅳ[수:로마십진]", "ⅳ[로마 숫자:수]"},
		"Ａ": {"Ａ[열:영숫자]", "Ａ[전각:로마자]"},
		"１": {"１[열:영숫자]", "１[전각:수]"},
	} {
		got := streamText(r, text)
		first, second := strings.Index(got, want[0]), strings.Index(got, want[1])
		if first < 0 || second < 0 || first > second {
			t.Errorf("%s: want %v in that order in\n%s", text, want, got)
		}
	}
	for text, want := range map[string]string{"ⅩⅢ": "ⅩⅢ[수:로마십진]", "Ⅷ세": "Ⅷ[수:로마십진]+세[단위:세]", "Ⅲ의": "Ⅲ[수:로마십진]+의[관형격:의]"} {
		if got := streamText(r, text); !strings.Contains(got, want) || strings.Contains(got, "[로마 숫자:수]") {
			t.Errorf("%s: want %s and no shape in\n%s", text, want, got)
		}
	}
}
