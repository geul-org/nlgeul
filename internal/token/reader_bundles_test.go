//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderBundles — the bundles printed again are the lines Lines gives; units are numbered in order; a contest, a branch of two words, a part written out, a part of two readings, a number and its unit, the signs of a pair, a path, a closing dot, a part with no trace in the letters, a unit of signs only, and a unit with no space whose readings all agree on one cut only (after 에서) come up as the structure says; a stretch holds every way its unit's readings read it (고양이가 five, the new word whole last)
package token

import (
	"reflect"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestReaderBundles(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	first := "작은 고양이가 물을 마신다."
	second := "우리는 사과 3개를 먹어보다."
	third := "고양이(물)을 ~/a/b를 했다."
	for _, s := range []string{first, second, third} {
		var lines []string
		for _, b := range r.Bundles(s) {
			lines = append(lines, render(b)...)
		}
		if want := r.Lines(s, false); !reflect.DeepEqual(lines, want) {
			t.Errorf("%s: printed again\n%q\nLines\n%q", s, lines, want)
		}
	}

	a := r.Bundles(first)
	var units []int
	var texts []string
	for _, b := range a {
		units, texts = append(units, b.Unit), append(texts, b.Text)
	}
	if want := []int{0, 1, 2, 3, 3}; !reflect.DeepEqual(units, want) {
		t.Fatalf("units %v, want %v", units, want)
	}
	if want := []string{"작은", "고양이가", "물을", "마신다", "."}; !reflect.DeepEqual(texts, want) {
		t.Fatalf("texts %q, want %q", texts, want)
	}
	if !a[0].Contest() || len(a[0].Branches) < 2 {
		t.Fatalf("작은 is not a contest: %+v", a[0])
	}
	adj := []Eojeol{{Parts: []Part{{Text: "작", Names: []Name{{"열쇠", "형용사"}}, Stem: true}, {Text: "은", Names: []Name{{"꾸밈", "은"}}}}, End: 2}}
	noun := []Eojeol{{Parts: []Part{{Text: "작", Names: []Name{{"열쇠", "체언"}}, Stem: true}, {Text: "은", Names: []Name{{"보조사", "은"}}}}, End: 2}}
	if !reflect.DeepEqual(a[0].Branches[0].Words, adj) || !reflect.DeepEqual(a[0].Branches[1].Words, noun) {
		t.Fatalf("작은: %+v", a[0].Branches)
	}
	subject := []Eojeol{{Parts: []Part{{Text: "고양이", Names: []Name{{"열쇠", "체언"}}, Stem: true}, {Text: "가", Names: []Name{{"주격", "가"}, {"보격", "가"}}}}, End: 4}}
	if !reflect.DeepEqual(a[1].Branches[0].Words, subject) || len(a[1].Branches) != 5 || a[1].Branches[4].String() != "고양이가" {
		t.Fatalf("고양이가: %+v", a[1].Branches)
	}
	drinks := []Eojeol{{Parts: []Part{{Text: "마시", Names: []Name{{"열쇠", "동사"}}, Stem: true}, {Text: "다", Hidden: "ㄴ", Names: []Name{{"맺음", "ㄴ다"}}, Ends: rule.KindFinal}}, End: 3}}
	if !reflect.DeepEqual(a[3].Branches[0].Words, drinks) || len(a[3].Branches) != 5 {
		t.Fatalf("마신다: %+v", a[3].Branches)
	}
	dot := []Eojeol{{Parts: []Part{{Text: ".", Names: []Name{{"마침표", "."}}, Stem: false}}, End: 3}}
	if len(a[4].Branches) != 1 || !reflect.DeepEqual(a[4].Branches[0].Words, dot) {
		t.Fatalf("closing dot: %+v", a[4].Branches)
	}

	c := r.Bundles(second)
	if len(c) != 5 || c[2].Text != "3개를" || c[3].Text != "먹어보다" || c[3].Unit != 3 {
		t.Fatalf("%s: %+v", second, c)
	}
	counted := []Eojeol{{Parts: []Part{{Text: "3", Names: []Name{{"수", ""}}, Stem: true}, {Text: "개", Names: []Name{{"단위", "개"}}}, {Text: "를", Names: []Name{{"목적격", "를"}}}}, End: 3}}
	if !reflect.DeepEqual(c[2].Branches[0].Words, counted) {
		t.Fatalf("3개를: %+v", c[2].Branches[0])
	}
	tried := []Eojeol{
		{Parts: []Part{{Text: "먹", Names: []Name{{"열쇠", "동사"}}, Stem: true}, {Text: "어", Names: []Name{{"맺음", "어"}}, Ends: rule.KindFinal | rule.KindLink}}, End: 2},
		{Parts: []Part{{Text: "보", Names: []Name{{"보조 용언", "보"}}, Stem: true}, {Text: "다", Names: []Name{{"맺음", "다"}}, Ends: rule.KindFinal}}, End: 4},
	}
	if !reflect.DeepEqual(c[3].Branches[0].Words, tried) {
		t.Fatalf("먹어보다: %+v", c[3].Branches[0])
	}

	d := r.Bundles(third)
	if len(d) != 7 || d[1].Text != "(" || d[3].Text != ")을" || d[4].Text != "~/a/b를" || d[5].Unit != 5 {
		t.Fatalf("%s: %+v", third, d)
	}
	opening := []Eojeol{{Parts: []Part{{Text: "(", Names: []Name{{"여는 묶음표", ""}}, Stem: true}}, End: 1}}
	if len(d[1].Branches) != 1 || !reflect.DeepEqual(d[1].Branches[0].Words, opening) {
		t.Fatalf("(: %+v", d[1].Branches)
	}
	closing := []Eojeol{{Parts: []Part{{Text: ")", Names: []Name{{"닫는 묶음표", ""}}, Stem: true}, {Text: "을", Names: []Name{{"목적격", "을"}}}}, End: 2}}
	if !reflect.DeepEqual(d[3].Branches[0].Words, closing) {
		t.Fatalf(")을: %+v", d[3].Branches[0])
	}
	path := []Eojeol{{Parts: []Part{{Text: "~/a/b", Names: []Name{{"경로", ""}}, Stem: true}, {Text: "를", Names: []Name{{"목적격", "를"}}}}, End: 6}}
	if len(d[4].Branches) != 2 || !reflect.DeepEqual(d[4].Branches[0].Words, path) && !reflect.DeepEqual(d[4].Branches[1].Words, path) {
		t.Fatalf("~/a/b를: %+v", d[4].Branches)
	}

	f := r.Bundles("학교다")
	copula := []Eojeol{{Parts: []Part{
		{Text: "학교", Names: []Name{{"열쇠", "체언"}}, Stem: true},
		{Text: "", Hidden: "이", Names: []Name{{"서술격", "이(빠짐)"}}, Unseen: true},
		{Text: "다", Names: []Name{{"맺음", "다"}}, Ends: rule.KindFinal},
	}, End: 3}}
	if len(f) < 1 || len(f[0].Branches) < 1 || !reflect.DeepEqual(f[0].Branches[0].Words, copula) {
		t.Fatalf("학교다: %+v", f)
	}

	g := r.Bundles("...")
	signs := []Eojeol{{Parts: []Part{{Text: "...", Stem: true}}}}
	if len(g) != 1 || g[0].Text != "..." || len(g[0].Branches) != 1 || !reflect.DeepEqual(g[0].Branches[0].Words, signs) || g[0].Contest() {
		t.Fatalf("signs only: %+v", g)
	}

	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	e := bare.Bundles("고양이가마당에서물을마신다")
	if len(e) != 2 || e[0].Text != "고양이가마당에서" || e[1].Text != "물을마신다" || e[0].Unit != 0 || e[1].Unit != 0 || !e[0].Contest() || !e[1].Contest() || len(e[1].Branches) != 4 {
		t.Fatalf("a unit cut where all its readings agree: %+v", e)
	}
}
