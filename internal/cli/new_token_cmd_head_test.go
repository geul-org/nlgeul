//ff:func feature=cli type=command control=sequence
//ff:what TestNewTokenCmdHead — --head limits the candidates printed for one contest (20 by default) and counts the rest; 0 prints all; a negative number is refused; --cuts and the trace lines are not cut
package cli

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestNewTokenCmdHead(t *testing.T) {
	t.Setenv("NLGEUL_WORDS", "")
	var out bytes.Buffer
	run := func(args ...string) error {
		out.Reset()
		root := newRootCmd()
		root.SetOut(&out)
		root.SetArgs(append([]string{"token"}, args...))
		return root.Execute()
	}
	lines := func() []string { return strings.Split(strings.TrimSuffix(out.String(), "\n"), "\n") }
	if err := run("--head", "0", "100m²이다"); err != nil || len(lines()) < 25 || strings.Contains(out.String(), "생략됨") {
		t.Fatalf("%v %q", err, out.String())
	}
	all := len(lines())
	if err := run("100m²이다"); err != nil || len(lines()) != 22 || lines()[21] != "  ("+strconv.Itoa(all-21)+"개 생략됨)" {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("--head", "5", "100m²이다"); err != nil || len(lines()) != 7 || lines()[6] != "  ("+strconv.Itoa(all-6)+"개 생략됨)" {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("--head", "-1", "100m²이다"); err == nil || !strings.Contains(err.Error(), "--head") {
		t.Fatal("a negative --head accepted", err)
	}
	if err := run("--head", "2", "--cuts", "100m²이다"); err != nil || strings.Contains(out.String(), "생략됨") {
		t.Fatalf("%v %q", err, out.String())
	}
}
