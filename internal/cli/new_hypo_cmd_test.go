//ff:func feature=cli type=command control=sequence
//ff:what TestNewHypoCmd — nlgeul hypo with no option prints what runHypo prints for page 0; a word table comes from --words or NLGEUL_WORDS, and it runs without one; a table that cannot be read, no sentence, two sentences and a sentence past the limit are refused; a --page below 1 is refused before the word table is read
package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/internal/hypo"
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
	if err := runHypo(&want, tabled, sentence, 0, hypo.PerPage); err != nil {
		t.Fatal(err)
	}
	bare, err := token.New(token.Options{})
	if err != nil {
		t.Fatal(err)
	}
	var wantBare bytes.Buffer
	if err := runHypo(&wantBare, bare, sentence, 0, hypo.PerPage); err != nil {
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
	if err := run("--words", "testdata/none.tsv", "--page", "0", sentence); err == nil || err.Error() != "--page는 1 이상이어야 한다: 0" || out.Len() != 0 {
		t.Fatalf("--page 0 with a missing word table: %v %q", err, out.String())
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
