//ff:func feature=cli type=command control=sequence
//ff:what TestNewParseCmdHead — --head limits the candidates printed in one box and counts the rest; 0 prints all; the default (20) does not cut the 6 candidates of 100m²이다, the same as 0; a negative number is refused; --cuts and the trace lines are not cut
package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewParseCmdHead(t *testing.T) {
	t.Setenv("NLGEUL_WORDS", "")
	var out bytes.Buffer
	run := func(args ...string) error {
		out.Reset()
		root := newRootCmd()
		root.SetOut(&out)
		root.SetArgs(append([]string{"parse"}, args...))
		return root.Execute()
	}
	if err := run("--head", "0", "100m²이다"); err != nil || strings.Contains(out.String(), "생략됨") {
		t.Fatalf("%v %q", err, out.String())
	}
	full := out.String()
	all := strings.Count(full, "\n")
	if err := run("100m²이다"); err != nil || out.String() != full {
		t.Fatalf("the default 20 cuts the 6 candidates: %v %q", err, out.String())
	}
	if err := run("--head", "3", "100m²이다"); err != nil || !strings.Contains(out.String(), "(3개 생략됨)") || strings.Count(out.String(), "\n") >= all {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("--head", "-1", "100m²이다"); err == nil || !strings.Contains(err.Error(), "--head") {
		t.Fatal("a negative --head accepted", err)
	}
	if err := run("--head", "2", "--cuts", "100m²이다"); err != nil || strings.Contains(out.String(), "생략됨") {
		t.Fatalf("%v %q", err, out.String())
	}
}
