//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderContestsMarks — the contests of a sentence with a plus, a middle dot and a quote after a digit: a plus between letters cuts the unit (one unit for each piece, the plus two candidates of the same verdict and preference), the middle dot is read as a word of its own beside the candidates that keep it in the word, a quote right after a digit is a contest of its own with the unit candidates beside the pair candidates, no contest holds the place name 기호, and Glued tells what stood without a space
package token

import (
	"slices"
	"testing"
)

func TestReaderContestsMarks(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	plus := r.Contests("곰+돌+밤 라")
	units := []int{}
	for _, b := range plus {
		units = append(units, b.Unit)
	}
	if len(plus) != 6 || units[0] != 0 || units[1] != 1 || units[2] != 2 || units[3] != 3 || units[4] != 4 || units[5] != 5 {
		t.Fatal("units", units, plus)
	}
	signs := plus[1]
	if signs.Text != "+" || len(signs.Branches) != 2 || signs.Branches[0].Words[0].Parts[0].Name.Kind != "더하기" || signs.Branches[1].Words[0].Parts[0].Name.Kind != "연결" || signs.Branches[0].Verdict != signs.Branches[1].Verdict || signs.Branches[0].Pref != signs.Branches[1].Pref {
		t.Fatal("the plus", signs)
	}
	for k, want := range []bool{false, true, true, true, true, false} {
		if plus[k].Glued != want {
			t.Errorf("plus glued %d: %v", k, plus[k].Glued)
		}
	}
	for _, text := range []string{"4+5", "곰+8"} {
		if got := r.Contests(text); len(got) != 3 || got[1].Text != "+" {
			t.Errorf("%q: %+v", text, got)
		}
	}
	for _, text := range []string{"C++", "i++", "b+=7", "+3", "~/a/b+c"} {
		if slices.Contains(textsOf(r.Contests(text)), "+") {
			t.Errorf("%q cut", text)
		}
	}
	dot := r.Contests("밥·국을")
	if len(dot) != 1 || len(dot[0].Branches) < 2 {
		t.Fatal(dot)
	}
	var apart, kept bool
	for _, br := range dot[0].Branches {
		apart = apart || len(br.Words) == 3 && br.Words[1].Parts[0].Name.Kind == "가운뎃점" && br.Words[2].Parts[len(br.Words[2].Parts)-1].Text == "을"
		kept = kept || len(br.Words) == 1 && br.Words[0].Parts[0].Text == "밥·국"
	}
	if !apart || !kept {
		t.Fatal("the dot as a word of its own and the word that keeps it", dot)
	}
	for _, text := range []string{"밥·국을", "밥ㆍ국", "밥・국 ·", "ㆍ밥", "곰+돌", "\"6\" 5'10\""} {
		if holdsMid(r.Contests(text)) {
			t.Errorf("%q holds a place name", text)
		}
	}
	quote := r.Contests("\"6\"")
	if len(quote) != 3 || quote[2].Text != "\"" || len(quote[2].Branches) != 3 || !quote[2].Glued {
		t.Fatal("a quote after a digit", quote)
	}
	kinds := []string{}
	for _, br := range quote[2].Branches {
		n := br.Words[0].Parts[0].Name
		kinds = append(kinds, n.Kind+":"+n.Form)
	}
	if kinds[0] != "닫기:1" || kinds[1] != "단위:인치" || kinds[2] != "단위:초" {
		t.Fatal(kinds)
	}
}
