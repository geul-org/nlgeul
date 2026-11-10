//ff:func feature=cli type=command control=sequence
//ff:what TestNewParseCmd — kopas parse prints a line per word; no sentence or two arguments are refused; the help names the three input conditions
package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewParseCmd(t *testing.T) {
	var out bytes.Buffer
	root := newRootCmd()
	root.SetOut(&out)
	root.SetArgs([]string{"parse", "그것이 문제였다"})
	if err := root.Execute(); err != nil || out.String() != "그것+이:주어\n문제+였+다:서술어\n" {
		t.Fatalf("%v %q", err, out.String())
	}
	root.SetArgs([]string{"parse"})
	if root.Execute() == nil {
		t.Fatal("no sentence accepted")
	}
	root.SetArgs([]string{"parse", "가", "나"})
	if root.Execute() == nil {
		t.Fatal("two arguments accepted")
	}
	out.Reset()
	root.SetArgs([]string{"parse", "--help"})
	if err := root.Execute(); err != nil || !strings.Contains(out.String(), "3. 조사를 반드시 올바로 붙인다.") {
		t.Fatalf("%v %q", err, out.String())
	}
}
