//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderLinesTypo — end to end with the table of testdata that has a misspelt form (맞기 for 맡기): the misspelt stem is printed as the right form with 오타정정 and the written form, below the same tail written right, with the readings it had before still there (the new word, the verb and the naming 기 with the particle); a bare 맞기 and 맞기가 have no such reading; the right form is printed as before; with no word table or a table of four columns there is none either
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesTypo(t *testing.T) {
	typoed, err := New(Options{Words: "testdata/words-typo.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	plain, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	const mark = "오타정정:맞기]"
	cases := []struct {
		sentence string
		has      []string
		lacks    []string
	}{
		{"맞겨", []string{"  맡기[열쇠:동사,오타정정:맞기]+(어)[맺음:어(합침)]\t성립 0.55", "  맞겨\t성립 0.33"}, nil},
		{"맞겼다", []string{"  맡기[열쇠:동사,오타정정:맞기]+(었)[선어말:ㅆ]+다[맺음:다]\t성립 0.55"}, nil},
		{"맞기고", []string{"  맡기[열쇠:동사,오타정정:맞기]+고[맺음:고]\t성립 0.55", "  맞기+고[맺음:고]\t성립 0.33"}, nil},
		{"맞기는", []string{"  맡기[열쇠:동사,오타정정:맞기]+는[꾸밈:는]\t성립 0.55", "  맞[열쇠:동사]+기[이름:기]+는[보조사:는]\t성립 1.00", "  맞기+는[보조사:는]\t성립 0.33"}, nil},
		{"맞긴다", []string{"  맡기[열쇠:동사,오타정정:맞기]+(ㄴ)다[맺음:ㄴ다]\t성립 0.55"}, nil},
		{"맞기면", []string{"  맡기[열쇠:동사,오타정정:맞기]+면[맺음:면]\t성립 0.55"}, nil},
		{"맞겨서", []string{"  맡기[열쇠:동사,오타정정:맞기]+(어)서[맺음:어서(합침)]\t성립 0.55"}, nil},
		{"맞겨줘", []string{"  맡기[열쇠:동사,오타정정:맞기]+(어)[맺음:어(합침)] 주[보조 용언:주]+(어)[맺음:어(합침)]\t성립 0.55"}, nil},
		{"일을 맞겨.", []string{"  맡기[열쇠:동사,오타정정:맞기]+(어)[맺음:어(합침)]\t성립 0.55"}, nil},
		{"맞기", []string{"  맞[열쇠:동사]+기[이름:기]\t성립 1.00"}, []string{mark}},
		{"맞기가", []string{"  맞[열쇠:동사]+기[이름:기]+가[주격:가|보격:가]\t성립 1.00"}, []string{mark}},
		{"맡겨", []string{"  맡기[열쇠:동사]+(어)[맺음:어(합침)]\t성립 0.98"}, []string{mark, "오타정정"}},
		{"맡기고", []string{"  맡기[열쇠:동사]+고[맺음:고]\t성립 0.98"}, []string{mark, "오타정정"}},
	}
	for _, c := range cases {
		got := strings.Join(typoed.Lines(c.sentence, false), "\n")
		for _, want := range c.has {
			if !strings.Contains(got, want) {
				t.Errorf("%s: missing %q in\n%s", c.sentence, want, got)
			}
		}
		for _, no := range c.lacks {
			if strings.Contains(got, no) {
				t.Errorf("%s: %q in\n%s", c.sentence, no, got)
			}
		}
	}
	for _, other := range []*Reader{plain, bare} {
		for _, s := range []string{"맞겨", "맞겼다", "맞기고", "맞기는", "맞긴다", "맞기면", "맞겨서", "맞겨줘", "맞기", "맞기가"} {
			if got := strings.Join(other.Lines(s, false), "\n"); strings.Contains(got, "오타정정") {
				t.Errorf("%s: no misspelt form is given, in\n%s", s, got)
			}
		}
	}
}
