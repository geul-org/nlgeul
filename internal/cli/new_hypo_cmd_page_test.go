//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestNewHypoCmdPage — nlgeul hypo with no option prints every hypothesis however many pages they fill; --page prints what runHypo prints for that page, of hypo.PerPage hypotheses or of --per-page; a --page that is not a number, a --page or a --per-page below 1 and a --per-page without --page are refused before anything is printed; a sentence past the limit is only counted and refused with --page too
package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/internal/hypo"
	"github.com/geul-org/nlgeul/internal/token"
)

func TestNewHypoCmdPage(t *testing.T) {
	t.Setenv("NLGEUL_WORDS", "")
	var out bytes.Buffer
	run := func(args ...string) error {
		out.Reset()
		root := newRootCmd()
		root.SetOut(&out)
		root.SetArgs(append([]string{"hypo"}, args...))
		return root.Execute()
	}
	bare, err := token.New(token.Options{})
	if err != nil {
		t.Fatal(err)
	}
	sentence := strings.TrimSpace(strings.Repeat("작은 ", 5))
	var all, second, third bytes.Buffer
	if runHypo(&all, bare, sentence, 0, hypo.PerPage) != nil || runHypo(&second, bare, sentence, 2, hypo.PerPage) != nil || runHypo(&third, bare, sentence, 3, 7) != nil {
		t.Fatal("runHypo refused the sentence")
	}
	if strings.Count(all.String(), "\n") <= 2*hypo.PerPage+1 || strings.Count(second.String(), "\n") != hypo.PerPage+1 || strings.Count(third.String(), "\n") != 7+1 {
		t.Fatalf("the sentence does not fill pages: %d, %d and %d lines", strings.Count(all.String(), "\n"), strings.Count(second.String(), "\n"), strings.Count(third.String(), "\n"))
	}

	if err := run(sentence); err != nil || out.String() != all.String() {
		t.Fatalf("no option: %v, %d lines, want %d", err, strings.Count(out.String(), "\n"), strings.Count(all.String(), "\n"))
	}
	if err := run("--page", "2", sentence); err != nil || out.String() != second.String() {
		t.Fatalf("--page 2: %v\n%q\nwant\n%q", err, out.String(), second.String())
	}
	if err := run("--page", "3", "--per-page", "7", sentence); err != nil || out.String() != third.String() {
		t.Fatalf("--page 3 --per-page 7: %v\n%q\nwant\n%q", err, out.String(), third.String())
	}

	for _, c := range []struct {
		args []string
		err  string // what the error says
	}{
		{[]string{"--page", "둘"}, `"둘" for "--page"`},
		{[]string{"--page", "0"}, "--page는 1 이상이어야 한다: 0"},
		{[]string{"--page=-2"}, "--page는 1 이상이어야 한다: -2"},
		{[]string{"--page", "1", "--per-page", "0"}, "--per-page는 1 이상이어야 한다: 0"},
		{[]string{"--page", "1", "--per-page=-4"}, "--per-page는 1 이상이어야 한다: -4"},
		{[]string{"--per-page", "7"}, "--per-page는 --page와 함께 준다"},
	} {
		if err := run(append(c.args, sentence)...); err == nil || !strings.Contains(err.Error(), c.err) || out.Len() != 0 {
			t.Errorf("%q: %v %q, want an error with %q and nothing printed", c.args, err, out.String(), c.err)
		}
	}

	if err := run("--page", "1", strings.TrimSpace(strings.Repeat("작은 ", 13))); err == nil || strings.Count(out.String(), "\n") != 1 || !strings.HasPrefix(out.String(), "가설 ") {
		t.Fatalf("past the limit with --page 1: %v %q", err, out.String())
	}
}
