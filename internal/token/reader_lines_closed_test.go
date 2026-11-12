//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderLinesClosed — where a full stop, a question mark or an exclamation mark closes a unit, the auxiliary 이야 and 야 right on a stem are not read (the copula reading and the vocative stay) and the vocative 야 is not read after a dependent noun; without such a sign, before dots in a row, after another particle and after a linking closing the auxiliary particle is read as before
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesClosed(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		has      []string
		lacks    []string
	}{
		{"마당이야.", []string{"마당[열쇠:체언]+이[서술격:이]+야[맺음:야]"}, []string{"[보조사:이야]"}},
		{"마당이야?", []string{"마당[열쇠:체언]+이[서술격:이]+야[맺음:야]"}, []string{"[보조사:이야]"}},
		{"마당이야", []string{"마당[열쇠:체언]+이야[보조사:이야]", "마당[열쇠:체언]+이[서술격:이]+야[맺음:야]"}, nil},
		{"마당이야...", []string{"마당[열쇠:체언]+이야[보조사:이야]"}, nil},
		{"마당이야 좋다.", []string{"마당[열쇠:체언]+이야[보조사:이야]"}, nil},
		{"문제야!", []string{"문제[열쇠:체언]+야[호격:야]", "문제[열쇠:체언]+(이)[서술격:이(빠짐)]+야[맺음:야]"}, []string{"보조사:야"}},
		{"문제야", []string{"문제[열쇠:체언]+야[보조사:야|호격:야]"}, nil},
		{"거야.", []string{"거[의존명사:거]+(이)[서술격:이(빠짐)]+야[맺음:야]"}, []string{"거[의존명사:거]+야["}},
		{"거야", []string{"거[의존명사:거]+야[보조사:야|호격:야]"}, nil},
		{"마당에서야.", []string{"마당[열쇠:체언]+에서[부사격:에서]+야[보조사:야|호격:야]"}, nil},
		{"먹고야.", []string{"먹[열쇠:동사]+고[맺음:고]+야[보조사:야]"}, nil},
	}
	for _, c := range cases {
		out := strings.Join(r.Lines(c.sentence, false), "\n")
		for _, want := range c.has {
			if !strings.Contains(out, want) {
				t.Errorf("%s: no reading %s in\n%s", c.sentence, want, out)
			}
		}
		for _, not := range c.lacks {
			if strings.Contains(out, not) {
				t.Errorf("%s: a reading with %s in\n%s", c.sentence, not, out)
			}
		}
	}
}
