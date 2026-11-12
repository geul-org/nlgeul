//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderContests — the contests printed again are the lines Lines gives; units are numbered in order; a contest of several candidates, a branch of two words, a part written out, a particle of two cases, one candidate each, a number and its unit, the signs of a pair, a path, a closing dot, a part with no trace in the letters, a unit of signs only (named, as ...), and a unit with no space whose candidates agree on no cut come up as the structure says, as one stretch holding every way the unit is read (none dropped for a limit), best first
package token

import (
	"reflect"
	"testing"
)

func TestReaderContests(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	first := "작은 고양이가 물을 마신다."
	second := "우리는 사과 3개를 먹어보다."
	third := "고양이(물)을 ~/a/b를 했다."
	for _, s := range []string{first, second, third} {
		var lines []string
		for _, b := range r.Contests(s) {
			lines = append(lines, render(b, 0)...)
		}
		if want := r.Lines(s, false); !reflect.DeepEqual(lines, want) {
			t.Errorf("%s: printed again\n%q\nLines\n%q", s, lines, want)
		}
	}

	a := r.Contests(first)
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
	if !a[0].Plural() || len(a[0].Branches) < 2 {
		t.Fatalf("작은 is not plural: %+v", a[0])
	}
	adj := []Eojeol{{Parts: []Part{{Text: "작", Name: Name{"열쇠", "형용사"}, Stem: true}, {Text: "은", Name: Name{"꾸밈", "은"}}}, End: 2}}
	noun := []Eojeol{{Parts: []Part{{Text: "작", Name: Name{"열쇠", "체언"}, Stem: true}, {Text: "은", Name: Name{"보조사", "은"}}}, End: 2}}
	if !reflect.DeepEqual(a[0].Branches[0].Words, adj) || !reflect.DeepEqual(a[0].Branches[1].Words, noun) {
		t.Fatalf("작은: %+v", a[0].Branches)
	}
	subject := []Eojeol{{Parts: []Part{{Text: "고양이", Name: Name{"열쇠", "체언"}, Stem: true}, {Text: "가", Name: Name{"주격", "가"}}}, End: 4}}
	complement := []Eojeol{{Parts: []Part{{Text: "고양이", Name: Name{"열쇠", "체언"}, Stem: true}, {Text: "가", Name: Name{"보격", "가"}}}, End: 4}}
	if n := len(a[1].Branches); !reflect.DeepEqual(a[1].Branches[0].Words, subject) || !reflect.DeepEqual(a[1].Branches[1].Words, complement) || n < 6 || a[1].Branches[n-1].String() != "고양이가" {
		t.Fatalf("고양이가: %+v", a[1].Branches)
	}
	drinks := []Eojeol{{Parts: []Part{{Text: "마시", Name: Name{"열쇠", "동사"}, Stem: true}, {Text: "다", Hidden: "ㄴ", Name: Name{"맺음", "ㄴ다"}}}, End: 3}}
	if !reflect.DeepEqual(a[3].Branches[0].Words, drinks) || len(a[3].Branches) != 5 {
		t.Fatalf("마신다: %+v", a[3].Branches)
	}
	dot := []Eojeol{{Parts: []Part{{Text: ".", Name: Name{"마침표", ""}, Stem: false}}, End: 3}}
	if len(a[4].Branches) != 1 || !reflect.DeepEqual(a[4].Branches[0].Words, dot) {
		t.Fatalf("closing dot: %+v", a[4].Branches)
	}

	c := r.Contests(second)
	if len(c) != 5 || c[2].Text != "3개를" || c[3].Text != "먹어보다" || c[3].Unit != 3 {
		t.Fatalf("%s: %+v", second, c)
	}
	counted := []Eojeol{{Parts: []Part{{Text: "3", Name: Name{"수", ""}, Stem: true}, {Text: "개", Name: Name{"단위", "개"}}, {Text: "를", Name: Name{"목적격", "를"}}}, End: 3}}
	if !reflect.DeepEqual(c[2].Branches[0].Words, counted) {
		t.Fatalf("3개를: %+v", c[2].Branches[0])
	}
	tried := []Eojeol{
		{Parts: []Part{{Text: "먹", Name: Name{"열쇠", "동사"}, Stem: true}, {Text: "어", Name: Name{"맺음", "어"}}}, End: 2},
		{Parts: []Part{{Text: "보", Name: Name{"보조 용언", "보"}, Stem: true}, {Text: "다", Name: Name{"맺음", "다"}}}, End: 4},
	}
	if !reflect.DeepEqual(c[3].Branches[0].Words, tried) {
		t.Fatalf("먹어보다: %+v", c[3].Branches[0])
	}

	d := r.Contests(third)
	if len(d) != 8 || d[1].Text != "(" || d[3].Text != ")" || d[4].Text != "을" || d[5].Text != "~/a/b를" || d[6].Unit != 5 {
		t.Fatalf("%s: %+v", third, d)
	}
	opening := []Eojeol{{Parts: []Part{{Text: "(", Name: Name{"열기", "1"}, Depth: 1, Stem: true}}, End: 1}}
	if len(d[1].Branches) != 1 || !reflect.DeepEqual(d[1].Branches[0].Words, opening) {
		t.Fatalf("(: %+v", d[1].Branches)
	}
	closing := []Eojeol{{Parts: []Part{{Text: ")", Name: Name{"닫기", "1"}, Depth: 1, Stem: true}}, End: 1}}
	if !reflect.DeepEqual(d[3].Branches[0].Words, closing) {
		t.Fatalf(")을: %+v", d[3].Branches[0])
	}
	onto := []Eojeol{{Parts: []Part{{Text: "을", Name: Name{"목적격", "을"}}}, End: 1, Attach: true}}
	if d[4].Unit != 3 || !d[4].Glued || !reflect.DeepEqual(d[4].Branches[0].Words, onto) || d[4].Branches[0].Verdict != 1 {
		t.Fatalf(")을: %+v", d[4])
	}
	path := []Eojeol{{Parts: []Part{{Text: "~/a/b", Name: Name{"경로", ""}, Stem: true}, {Text: "를", Name: Name{"목적격", "를"}}}, End: 6}}
	if len(d[5].Branches) != 2 || !reflect.DeepEqual(d[5].Branches[0].Words, path) && !reflect.DeepEqual(d[5].Branches[1].Words, path) {
		t.Fatalf("~/a/b를: %+v", d[5].Branches)
	}

	f := r.Contests("학교다")
	copula := []Eojeol{{Parts: []Part{
		{Text: "학교", Name: Name{"열쇠", "체언"}, Stem: true},
		{Text: "", Hidden: "이", Name: Name{"서술격", "이(빠짐)"}, Unseen: true},
		{Text: "다", Name: Name{"맺음", "다"}},
	}, End: 3}}
	if len(f) < 1 || len(f[0].Branches) < 1 || !reflect.DeepEqual(f[0].Branches[0].Words, copula) {
		t.Fatalf("학교다: %+v", f)
	}

	g := r.Contests("...")
	signs := []Eojeol{{Parts: []Part{{Text: "...", Name: Name{"줄임표", ""}}}}}
	if len(g) != 1 || g[0].Text != "..." || len(g[0].Branches) != 1 || !reflect.DeepEqual(g[0].Branches[0].Words, signs) || g[0].Plural() {
		t.Fatalf("signs only: %+v", g)
	}

	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	e := bare.Contests("고양이가마당에서물을마신다")
	right := "고양이+가[주격:가] 마당+에서[주격:에서] 물+을[목적격:을] 마신+다[맺음:다]"
	if len(e) != 1 || e[0].Text != "고양이가마당에서물을마신다" || e[0].Unit != 0 || !e[0].Plural() || len(e[0].Branches) < 4 || !hasBranch(e[0].Branches, right) || !branchesInOrder(e[0].Branches) {
		t.Fatalf("a unit whose candidates agree on no cut is one stretch: %+v", e)
	}
}
