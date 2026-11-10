//ff:func feature=cli type=command control=sequence
//ff:what TestTokenInput — one argument is the sentence; with --stdin and no argument every line is a sentence; neither or both is refused
package cli

import (
	"strings"
	"testing"
)

func TestTokenInput(t *testing.T) {
	if got, err := tokenInput([]string{"가 나"}, false, nil); err != nil || len(got) != 1 {
		t.Fatal(got, err)
	}
	if got, err := tokenInput(nil, true, strings.NewReader("가\n나 다\n")); err != nil || len(got) != 2 || got[1] != "나 다" {
		t.Fatal(got, err)
	}
	if _, err := tokenInput(nil, false, nil); err == nil {
		t.Fatal("no sentence accepted")
	}
	if _, err := tokenInput([]string{"가"}, true, strings.NewReader("")); err == nil {
		t.Fatal("an argument with --stdin accepted")
	}
}
