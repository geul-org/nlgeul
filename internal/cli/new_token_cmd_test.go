//ff:func feature=cli type=command control=sequence
//ff:what TestNewTokenCmd — kopas token prints the words of a unit; --cuts word ends; --trace the judgments; --stdin reads lines; a word table comes from --words or KOPAS_WORDS; no sentence and a table that cannot be read are refused
package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewTokenCmd(t *testing.T) {
	t.Setenv("KOPAS_WORDS", "")
	var out bytes.Buffer
	run := func(in string, args ...string) error {
		out.Reset()
		root := newRootCmd()
		root.SetOut(&out)
		root.SetIn(strings.NewReader(in))
		root.SetArgs(append([]string{"token"}, args...))
		return root.Execute()
	}
	if err := run("", "아버지가"); err != nil || !strings.Contains(out.String(), "\n  아버지+가[주격:가|보격:가]\t성립 0.33") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--cuts", "아버지가"); err != nil || !strings.HasPrefix(out.String(), "아버지가\t4;") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--trace", "아버지가"); err != nil || !strings.Contains(out.String(), "주격:가\t1.00\t통과") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("아버지가\n아버지가\n", "--stdin", "--cuts"); err != nil || strings.Count(out.String(), "아버지가\t4;") != 2 {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--words", "testdata/words.tsv", "마당에서"); err != nil || !strings.Contains(out.String(), "\n  마당[열쇠:체언]+에서[주격:에서|부사격:에서]\t성립 1.00") {
		t.Fatalf("%v %q", err, out.String())
	}
	t.Setenv("KOPAS_WORDS", "testdata/words.tsv")
	if err := run("", "--layers", "1", "마당에서"); err != nil || !strings.Contains(out.String(), "마당[열쇠:체언]+에서") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--words", "testdata/none.tsv", "마당에서"); err == nil {
		t.Fatal("a missing word table accepted")
	}
	if err := run(""); err == nil {
		t.Fatal("no sentence accepted")
	}
}
