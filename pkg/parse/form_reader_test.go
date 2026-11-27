//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestFormReader — a letter of the tables of special letters is named by its shape and what it writes: a number of one letter of any shape with a particle or a unit after it, as a digit is (⓪, ⒈, ❶, ½, ②번); a sign of the tables at a unit's edge is a contest of its own with that name (ⓐ, ㉠, ㉮, ⒜, ？, each letter of ⓐⓑ); a letter of the tables that is no sign and writes no number is an atom with that name (㊀); a unit of the unit table written with such letters stays a unit after its number (5m², 30㎡)
package parse

import (
	"strings"
	"testing"
)

func TestFormReader(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for text, want := range map[string]string{
		"⓪":   "⓪[원:수]",
		"⒈":   "⒈[점 붙음:수]",
		"❶":   "❶[검은 원:수]",
		"⓵":   "⓵[겹원:수]",
		"½":   "½[분수:수]",
		"❷의":  "❷[검은 원:수]+의[관형격:의]",
		"⓪번":  "⓪[원:수]+번[단위:번]",
		"ⓐ":   "ⓐ[원:로마자]",
		"㉠":   "㉠[원:한글]",
		"㉮":   "㉮[원:한글]",
		"⒜":   "⒜[괄호:로마자]",
		"？":   "？[전각:기호]",
		"ⓐⓑ":  "ⓑ[원:로마자]",
		"사과ⓐ": "ⓐ[원:로마자]",
		"㊀":   "㊀[원:한자]",
		"㊀의":  "㊀[원:한자]+의[관형격:의]",
		"5m²": "5[수:십진]+m²[단위:제곱미터]",
		"30㎡": "30[수:십진]+㎡[단위:제곱미터]",
		"㈎":   "㈎[괄호:한글]",
		"㈜사과": "㈜[괄호:한글]",
	} {
		if got := streamText(r, text); !strings.Contains(got, want) {
			t.Errorf("%s: want %s in\n%s", text, want, got)
		}
	}
	for _, text := range []string{"ⓐ", "㉠", "㊀", "？"} {
		if got := streamText(r, text); strings.Contains(got, "[수") {
			t.Errorf("%s is no number: %s", text, got)
		}
	}
}
