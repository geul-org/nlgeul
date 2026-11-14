//ff:func feature=cli type=command control=sequence
//ff:what TestNewTokenCmdBox — nlgeul token draws boxes by default (a title, the right line of every box of the contest of ways in one cell, the line "1등 길:"), --format plain gives the old lines, --head counts the ways left out as well as the candidates, --trace puts the lines of a unit before its boxes, --stdin parts the sentences by an empty line, --cuts stays as it was
package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewTokenCmdBox(t *testing.T) {
	t.Setenv("NLGEUL_WORDS", "")
	var out bytes.Buffer
	run := func(in string, args ...string) error {
		out.Reset()
		root := newRootCmd()
		root.SetOut(&out)
		root.SetIn(strings.NewReader(in))
		root.SetArgs(append([]string{"token"}, args...))
		return root.Execute()
	}
	if err := run("", "아버지가"); err != nil || !strings.HasPrefix(out.String(), "┌ [0-4] 아버지가 ─ 후보경합 ") || !strings.Contains(out.String(), "\n│ 아버지+가[주격:가]") || !strings.Contains(out.String(), "\n1등 길: 아버지가  (λ 합 ") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--head", "2", "--words", "testdata/words.tsv", "고양이가마당에서"); err != nil || !strings.Contains(out.String(), "어절경합 [0-8]") || !strings.Contains(out.String(), "개 생략됨)") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--trace", "물을"); err != nil || strings.HasPrefix(out.String(), "┌") || !strings.Contains(out.String(), "통과") || strings.Index(out.String(), "┌") < strings.Index(out.String(), "통과") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("물을\n물을\n", "--stdin"); err != nil || strings.Count(out.String(), "1등 길: 물을") != 2 || !strings.Contains(out.String(), "\n\n┌") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--format", "plain", "--trace", "물을"); err != nil || strings.Contains(out.String(), "┌") || !strings.Contains(out.String(), "\t경합(") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--cuts", "물을"); err != nil || strings.Contains(out.String(), "┌") || !strings.HasPrefix(out.String(), "물을\t") {
		t.Fatalf("%v %q", err, out.String())
	}
}
