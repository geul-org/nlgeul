//ff:func feature=cli type=command control=sequence
//ff:what TestNewTokenCmdFormat — --format box is the default output as it is, --format json is one line of JSON for each sentence (numbered from 0, the Korean letters not escaped, a new word is not known), a value other than box and json and json with --format plain, --cuts or --marks are refused, and --format plain is the old lines, --plain is no longer an option
package cli

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestNewTokenCmdFormat(t *testing.T) {
	t.Setenv("NLGEUL_WORDS", "")
	var out bytes.Buffer
	run := func(in string, args ...string) error {
		out.Reset()
		root := newRootCmd()
		root.SetOut(&out)
		root.SetErr(&bytes.Buffer{})
		root.SetIn(strings.NewReader(in))
		root.SetArgs(append([]string{"token"}, args...))
		return root.Execute()
	}
	if err := run("", "아버지가"); err != nil {
		t.Fatal(err)
	}
	box := out.String()
	if err := run("", "--format", "box", "아버지가"); err != nil || out.String() != box {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--format", "plain", "물을"); err != nil || !strings.Contains(out.String(), "\t경합(") {
		t.Fatalf("%v %q", err, out.String())
	}
	if err := run("", "--plain", "물을"); err == nil {
		t.Fatal("--plain stays")
	}
	if err := run("쪼간만 해줘\n물을\n", "--format", "json", "--stdin", "--head", "1"); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(out.String(), "\n"), "\n")
	if len(lines) != 2 || !strings.Contains(lines[0], `"text":"쪼간만 해줘"`) || strings.Contains(out.String(), `\u`) {
		t.Fatalf("%q", out.String())
	}
	var d struct {
		Sentence int `json:"sentence"`
		Segments []struct {
			Words []struct {
				Text       string `json:"text"`
				Candidates []struct {
					Parts []struct {
						Text  string  `json:"text"`
						Name  *string `json:"name"`
						Known bool    `json:"known"`
					} `json:"parts"`
				} `json:"candidates"`
			} `json:"words"`
		} `json:"segments"`
	}
	if err := json.Unmarshal([]byte(lines[0]), &d); err != nil || d.Sentence != 0 || len(d.Segments) != 2 {
		t.Fatalf("%v %q", err, lines[0])
	}
	stem := d.Segments[0].Words[0].Candidates[0].Parts[0]
	whole := d.Segments[1].Words[0].Candidates[0].Parts[0]
	if stem.Text != "쪼간" || stem.Known || stem.Name != nil || whole.Text != "해줘" || whole.Known || len(d.Segments[0].Words[0].Candidates) < 2 {
		t.Fatalf("%v %v", stem, whole)
	}
	if err := json.Unmarshal([]byte(lines[1]), &d); err != nil || d.Sentence != 1 {
		t.Fatalf("%v %q", err, lines[1])
	}
	refused := func(args ...string) bool { return run("", args...) != nil }
	if !refused("--format", "xml", "물을") || !refused("--format", "json", "--cuts", "물을") || !refused("--format", "json", "--marks", "물을") {
		t.Fatal("not refused")
	}
}
