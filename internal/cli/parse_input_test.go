//ff:func feature=cli type=command control=sequence
//ff:what TestParseInput — one argument is the sentence; with --stdin and no argument every line is a sentence; neither or both is refused
package cli

import (
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	if got, err := parseInput([]string{"가 나"}, false, nil); err != nil || len(got) != 1 {
		t.Fatal(got, err)
	}
	if got, err := parseInput(nil, true, strings.NewReader("가\n나 다\n")); err != nil || len(got) != 2 || got[1] != "나 다" {
		t.Fatal(got, err)
	}
	if _, err := parseInput(nil, false, nil); err == nil {
		t.Fatal("no sentence accepted")
	}
	if _, err := parseInput([]string{"가"}, true, strings.NewReader("")); err == nil {
		t.Fatal("an argument with --stdin accepted")
	}
}
