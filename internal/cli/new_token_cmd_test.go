//ff:func feature=cli type=command control=sequence
//ff:what TestNewTokenCmd — kopas token prints the words of a unit; --cuts prints word ends; --trace the judgments; no sentence is refused
package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewTokenCmd(t *testing.T) {
	var out bytes.Buffer
	root := newRootCmd()
	root.SetOut(&out)
	root.SetArgs([]string{"token", "아버지가"})
	if err := root.Execute(); err != nil || !strings.Contains(out.String(), "\n  아버지+가[주격:가|보격:가]\t성립 0.33") {
		t.Fatalf("%v %q", err, out.String())
	}
	out.Reset()
	root.SetArgs([]string{"token", "--cuts", "아버지가"})
	if err := root.Execute(); err != nil || !strings.HasPrefix(out.String(), "아버지가\t4;") {
		t.Fatalf("%v %q", err, out.String())
	}
	out.Reset()
	root.SetArgs([]string{"token", "--cuts=false", "--trace", "아버지가"})
	if err := root.Execute(); err != nil || !strings.Contains(out.String(), "주격:가\t1.00\t통과") {
		t.Fatalf("%v %q", err, out.String())
	}
	root.SetArgs([]string{"token"})
	if root.Execute() == nil {
		t.Fatal("no sentence accepted")
	}
}
