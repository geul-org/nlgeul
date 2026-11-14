//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderLinesHangulBound — end to end, with and without the word table of testdata: after a run of Latin letters or digits, a number, a span of code, a path or a file name the closings 은, 는, 을, 아, 어, 나 and 여 are not read and the particle (or the vocative) stays; the closings 다, 고, 면, 지, 네 after such a run, and those closings after a Hangul syllable, are read as before, and so are 하 and the copula's stem between such a run and a closing; with the word table the closings after a verb or adjective stem stand
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestReaderLinesHangulBound(t *testing.T) {
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	lexed, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		has      []string
		lacks    []string
		words    bool
	}{
		{"push는", []string{"push[열:영숫자]+는[보조사:는]"}, []string{"[꾸밈:는]"}, false},
		{"API을", []string{"API[열:영숫자]+을[목적격:을]"}, []string{"[꾸밈:을]"}, false},
		{"abc은", []string{"abc[열:영숫자]+은[보조사:은]"}, []string{"[꾸밈:은]"}, false},
		{"v2나", []string{"v2[열:영숫자]+나[보조사:나]"}, []string{"[맺음:나]"}, false},
		{"`run`는", []string{"`run`[코드]+는[보조사:는]"}, []string{"[꾸밈:는]"}, false},
		{"a/b어", []string{"a/b[경로]"}, []string{"[맺음:어]"}, false},
		{"task.py여", []string{"task.py[파일명]+여[호격:여]", "task.py[도메인]+여[호격:여]"}, []string{"[맺음:여]"}, false},
		{"3는", []string{"3[수]+는[보조사:는]"}, []string{"[꾸밈:는]"}, false},
		{"OK아", []string{"OK[열:영숫자]+아[호격:아]"}, []string{"[맺음:아]"}, false},
		{"abc다", []string{"abc[열:영숫자]+다[맺음:다]"}, nil, false},
		{"abc고", []string{"abc[열:영숫자]+고[맺음:고]"}, nil, false},
		{"abc네", []string{"abc[열:영숫자]+네[맺음:네]"}, nil, false},
		{"API인", []string{"API[열:영숫자]+이[서술격:이]+(ㄴ)[꾸밈:ㄴ]"}, nil, false},
		{"3개는", []string{"3[수]+개[단위:개]+는[보조사:는]"}, nil, false},
		{"먹는", []string{"먹[열쇠:동사]+는[꾸밈:는]"}, nil, true},
		{"작은", []string{"작[열쇠:형용사]+은[꾸밈:은]"}, nil, true},
		{"먹을", []string{"먹[열쇠:동사]+을[꾸밈:을]"}, nil, true},
		{"살아", []string{"살[열쇠:동사]+아[맺음:아]"}, nil, true},
		{"먹어", []string{"먹[열쇠:동사]+어[맺음:어]"}, nil, true},
		{"가나", []string{"가[열쇠:동사]+나[맺음:나]"}, nil, true},
		{"하여", []string{"하[열쇠:동사]+여[맺음:여]"}, nil, true},
		{"push하는", []string{"push[열:영숫자] 하[열쇠:동사]+는[꾸밈:는]"}, nil, true},
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
