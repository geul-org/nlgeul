//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderLinesEmph — end to end, with and without the word table of testdata: the emphatic 가 stands only after an adverbial particle (학교에서가, 마당에가); after a number's unit, after an honorific, after the naming closing 기 and after a closing that asks it is not read, and the case particle 가 stays there as it was
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestReaderLinesEmph(t *testing.T) {
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	lexed, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	const emph = "가[보조사:가(강조)]"
	cases := []struct {
		sentence string
		has      []string
		lacks    []string
		words    bool
	}{
		{"집에서가", []string{"집+에서[부사격:에서]+" + emph}, nil, false},
		{"집에가", []string{"집+에[부사격:에]+" + emph}, nil, false},
		{"학교에서가", []string{"학교[열쇠:체언]+에서[부사격:에서]+" + emph}, nil, true},
		{"마당에가", []string{"마당[열쇠:체언]+에[부사격:에]+" + emph}, nil, true},
		{"3건가", nil, []string{emph}, false},
		{"3건이", []string{"3[수]+건[단위:건]+이["}, nil, false},
		{"7개가", []string{"7[수]+개[단위:개]+가[주격:가]"}, []string{emph}, false},
		{"두개가", []string{"두[수]+개[단위:개]+가[주격:가]"}, []string{emph}, true},
		{"가심가", nil, []string{emph}, true},
		{"먹기가", []string{"먹[열쇠:동사]+기[이름:기]+가[주격:가]"}, []string{emph}, true},
		{"먹는지가", []string{"먹[열쇠:동사]+는지[맺음:는지]+가[주격:가]"}, []string{emph}, true},
		{"살지가", []string{"+가[주격:가]"}, []string{emph}, true},
	}
	for _, c := range cases {
		readers := []*Reader{lexed}
		if !c.words {
			readers = append(readers, bare)
		}
		for _, r := range readers {
			out := strings.Join(r.Lines(c.sentence, false), "\n")
			if i := slices.IndexFunc(c.has, func(want string) bool { return !strings.Contains(out, want) }); i >= 0 {
				t.Errorf("%s: no %q\n%s", c.sentence, c.has[i], out)
			}
			if i := slices.IndexFunc(c.lacks, func(not string) bool { return strings.Contains(out, not) }); i >= 0 {
				t.Errorf("%s: %q read\n%s", c.sentence, c.lacks[i], out)
			}
		}
	}
}
