//ff:func feature=cli type=command control=sequence
//ff:what TestNewTokenCmd — nlgeul token prints the words of a unit; --format plain the old lines; the boxes by default; --cuts word ends; --marks word ends with the spacing rules (--spacing: all by default, none, or names); --trace the judgments; --stdin reads lines; a word table comes from --words or NLGEUL_WORDS; no sentence, a negative --head and a table that cannot be read are refused
package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewTokenCmd(t *testing.T) {
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
	if err := run("", "--format", "plain", "아버지가"); err != nil || !strings.Contains(out.String(), "\n  아버지+가[주격:가]\t성립 0.33") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--cuts", "아버지가"); err != nil || !strings.HasPrefix(out.String(), "아버지가\t4;") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--marks", "아버지가"); err != nil || !strings.HasPrefix(out.String(), "아버지가\t") || strings.Contains(out.String(), ";\n") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--marks", "--words", "testdata/words.tsv", "고양이가마당에서"); err != nil || !strings.Contains(out.String(), "3:2/41") || !strings.Contains(out.String(), "3:2/49/50") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--marks", "--spacing", "none", "--words", "testdata/words.tsv", "고양이가마당에서"); err != nil || strings.Contains(out.String(), "/") || !strings.Contains(out.String(), "3:-") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--marks", "--spacing", "41,50", "--words", "testdata/words.tsv", "고양이가마당에서"); err != nil || strings.Contains(out.String(), "2/") || !strings.Contains(out.String(), "3:41") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--spacing", "99", "아버지가"); err == nil || !strings.Contains(err.Error(), "99") {
		t.Fatal("a spacing rule that is not one accepted", err)
	}
	if err := run("", "--format", "xml", "아버지가"); err == nil || !strings.Contains(err.Error(), "xml") {
		t.Fatal("a format that is not one accepted", err)
	}
	if err := run("", "--format", "json", "--cuts", "아버지가"); err == nil {
		t.Fatal("json with cuts accepted")
	}
	if err := run("", "--format", "json", "아버지가"); err != nil || !strings.HasPrefix(out.String(), `{"sentence":0,"text":"아버지가"`) {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--trace", "아버지가"); err != nil || !strings.Contains(out.String(), "주격:가\t1.00\t통과") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("아버지가\n아버지가\n", "--stdin", "--cuts"); err != nil || strings.Count(out.String(), "아버지가\t4;") != 2 {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--format", "plain", "--words", "testdata/words.tsv", "마당에서"); err != nil || !strings.Contains(out.String(), "\n  마당[열쇠:체언]+에서[주격:에서]\t성립 ") {
		t.Fatalf("%v %q", err, out.String())
	}
	t.Setenv("NLGEUL_WORDS", "testdata/words.tsv")
	if err := run("", "--layers", "1", "마당에서"); err != nil || !strings.Contains(out.String(), "마당[열쇠:체언]+에서") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--words", "testdata/none.tsv", "마당에서"); err == nil {
		t.Fatal("a missing word table accepted")
	}
	if err := run(""); err == nil {
		t.Fatal("no sentence accepted")
	}
	if err := run("", "--head", "-3", "아버지가"); err == nil || !strings.Contains(err.Error(), "--head") {
		t.Fatal("a negative --head accepted", err)
	}
}
