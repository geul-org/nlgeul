//ff:func feature=cli type=command control=sequence
//ff:what TestNewHypoCmd — nlgeul hypo prints what runHypo prints; a word table comes from --words or NLGEUL_WORDS, and it runs without one; a table that cannot be read, no sentence, two sentences and a sentence past the limit are refused
package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/internal/token"
)

func TestNewHypoCmd(t *testing.T) {
	t.Setenv("NLGEUL_WORDS", "")
	var out bytes.Buffer
	run := func(args ...string) error {
		out.Reset()
		root := newRootCmd()
		root.SetOut(&out)
		root.SetArgs(append([]string{"hypo"}, args...))
		return root.Execute()
	}
	sentence := "우리는 마당에서 사과를 먹었다."
	tabled, err := token.New(token.Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	var want bytes.Buffer
	if err := runHypo(&want, tabled, sentence); err != nil {
		t.Fatal(err)
	}
	bare, err := token.New(token.Options{})
	if err != nil {
		t.Fatal(err)
	}
	var wantBare bytes.Buffer
	if err := runHypo(&wantBare, bare, sentence); err != nil {
		t.Fatal(err)
	}
	if want.String() == wantBare.String() || !strings.Contains(want.String(), "[열쇠:") || strings.Contains(wantBare.String(), "[열쇠:") {
		t.Fatalf("the word table makes no difference:\n%q\n%q", want.String(), wantBare.String())
	}

	if err := run("--words", "testdata/words.tsv", sentence); err != nil || out.String() != want.String() {
		t.Fatalf("--words: %v\n%q\nwant\n%q", err, out.String(), want.String())
	}
	if err := run(sentence); err != nil || out.String() != wantBare.String() {
		t.Fatalf("no word table: %v\n%q\nwant\n%q", err, out.String(), wantBare.String())
	}
	t.Setenv("NLGEUL_WORDS", "testdata/words.tsv")
	if err := run(sentence); err != nil || out.String() != want.String() {
		t.Fatalf("NLGEUL_WORDS: %v\n%q\nwant\n%q", err, out.String(), want.String())
	}
	t.Setenv("NLGEUL_WORDS", "")
	if err := run("--words", "testdata/none.tsv", sentence); err == nil {
		t.Fatal("a missing word table accepted")
	}
	if err := run(); err == nil {
		t.Fatal("no sentence accepted")
	}
	if err := run(sentence, sentence); err == nil {
		t.Fatal("two sentences accepted")
	}
	if err := run(strings.TrimSpace(strings.Repeat("작은 ", 13))); err == nil || strings.Count(out.String(), "\n") != 1 || !strings.HasPrefix(out.String(), "가설 ") {
		t.Fatalf("past the limit: %v %q", err, out.String())
	}
}
