//ff:func feature=token type=reader control=sequence
//ff:what TestCuts — a line for each unit with the word ends of its readings, best first
package token

import (
	"strings"
	"testing"
)

func TestCuts(t *testing.T) {
	got := Cuts("아버지가 다음으로써")
	if len(got) != 2 || !strings.HasPrefix(got[0], "아버지가\t4;") || !strings.HasPrefix(got[1], "다음으로써\t5;4,5") {
		t.Fatal(got)
	}
	if Cuts("") != nil {
		t.Fatal("empty")
	}
}
