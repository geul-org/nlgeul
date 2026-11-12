//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderBundlesSigns — the signs of a sentence as bundles: a nested bracket reading has the names 열기/닫기 with the depth as the form and as Part.Depth, a sign read both as a pair and as its letter has the letter as a part with no name and depth 0, # heads a line as 제목 with its count as the depth, the signs are a bundle of their own in the unit of the word they stand at, and the bundles printed again are the lines Lines gives
package token

import (
	"reflect"
	"testing"
)

func TestReaderBundlesSigns(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"함수(인자(안쪽))를 호출", "'I'm (O'Brien)'", "## 제목2", "#태그 @홍길동에게", "밥 먹었어?!"} {
		var lines []string
		for _, b := range r.Bundles(s) {
			lines = append(lines, render(b, 0)...)
		}
		if want := r.Lines(s, false); !reflect.DeepEqual(lines, want) {
			t.Errorf("%s: printed again\n%q\nLines\n%q", s, lines, want)
		}
	}

	nested := r.Bundles("함수(인자(안쪽))를")
	var texts []string
	var units []int
	for _, b := range nested {
		texts, units = append(texts, b.Text), append(units, b.Unit)
	}
	if want := []string{"함수", "(", "인자", "(", "안쪽", ")", ")", "를"}; !reflect.DeepEqual(texts, want) {
		t.Fatalf("texts %q, want %q", texts, want)
	}
	if want := []int{0, 1, 2, 3, 4, 5, 6, 6}; !reflect.DeepEqual(units, want) {
		t.Fatalf("units %v, want %v", units, want)
	}
	part := func(b Bundle, branch int) Part { return b.Branches[branch].Words[0].Parts[0] }
	if p := part(nested[1], 0); p.Text != "(" || p.Name != (Name{"열기", "1"}) || p.Depth != 1 || !p.Stem {
		t.Fatalf("(: %+v", p)
	}
	if p := part(nested[3], 0); p.Name != (Name{"열기", "2"}) || p.Depth != 2 {
		t.Fatalf("inner (: %+v", p)
	}
	if p := part(nested[5], 0); p.Name != (Name{"닫기", "2"}) || p.Depth != 2 {
		t.Fatalf("inner ): %+v", p)
	}
	closing := nested[6]
	if p := part(closing, 0); p.Name != (Name{"닫기", "1"}) || p.Depth != 1 || p.Text != ")" {
		t.Fatalf(")를: %+v", p)
	}
	if len(closing.Branches) != 1 || len(closing.Branches[0].Words[0].Parts) != 1 {
		t.Fatalf("a closing sign is a bundle of its own: %+v", closing)
	}
	onto := nested[7]
	if w := onto.Branches[0].Words[0]; onto.Unit != 6 || !onto.Glued || !w.Attach || len(w.Parts) != 1 || w.Parts[0].Text != "를" || w.Parts[0].Stem || onto.Branches[0].Verdict != 1 {
		t.Fatalf("a particle written onto the closing sign is a bundle of its own, with no stem: %+v", onto)
	}
	for k, b := range nested {
		if b.Glued != (k > 0) {
			t.Fatalf("glued %d: %+v", k, b)
		}
	}

	quote := r.Bundles("'I'm (O'Brien)'")
	if len(quote) != 6 || quote[0].Text != "'" || len(quote[0].Branches) != 2 {
		t.Fatalf("'I'm (O'Brien)': %+v", quote)
	}
	if p := part(quote[0], 0); p.Name != (Name{"열기", "1"}) || p.Depth != 1 {
		t.Fatalf("the quote as a pair: %+v", p)
	}
	if p := part(quote[0], 1); p.Name != (Name{}) || p.Depth != 0 || p.Text != "'" || p.Label() != "'" {
		t.Fatalf("the quote as its letter: %+v", p)
	}
	if last := quote[5]; len(last.Branches) != 3 || last.Branches[0].Words[0].Parts[0].Depth != 1 || last.Branches[1].Words[0].Parts[0].Name.Kind != "닫기" {
		t.Fatalf("the last quote: %+v", last)
	}
	paren := quote[2]
	if len(paren.Branches) != 2 || part(paren, 0).Depth != 1 || part(paren, 1).Depth != 2 {
		t.Fatalf("a bracket in a quote has two depths: %+v", paren)
	}

	heading := r.Bundles("## 제목2")
	if len(heading) != 2 || heading[0].Text != "##" || len(heading[0].Branches) != 2 {
		t.Fatalf("## 제목2: %+v", heading)
	}
	if p := part(heading[0], 0); p.Name != (Name{"제목2", ""}) || p.Depth != 2 {
		t.Fatalf("heading: %+v", p)
	}
	if p := part(heading[0], 1); p.Name != (Name{}) || !p.Stem || p.Depth != 0 {
		t.Fatalf("heading as its letters: %+v", p)
	}

	ask := r.Bundles("밥 먹었어?!")
	if len(ask) != 4 || ask[2].Text != "?" || ask[3].Text != "!" || ask[2].Unit != 1 || ask[3].Unit != 1 || part(ask[2], 0).Name.Kind != "물음표" || part(ask[3], 0).Name.Kind != "느낌표" {
		t.Fatalf("?!: %+v", ask)
	}
}
