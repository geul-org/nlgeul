//ff:func feature=parse type=helper control=sequence
//ff:what TestContentNames — the names of the lines of one content, each once and in the order of the lines; lines of another content are left out, and the tables held name 원:수 and 괄호:수 among the letters that write a number
package parse

import (
	"slices"
	"testing"
)

func TestContentNames(t *testing.T) {
	list := []form{{'①', "원", "수", "1", "1", ""}, {'ⓐ', "원", "로마자", "a", "", ""}, {'②', "원", "수", "2", "2", ""}, {'⑴', "괄호", "수", "1", "1", ""}}
	if got := contentNames(list, "수"); !slices.Equal(got, []string{"원:수", "괄호:수"}) {
		t.Fatal(got)
	}
	if got := contentNames(list, "한글"); got != nil {
		t.Fatal(got)
	}
	if !slices.Contains(formNumberNames, "원:수") || !slices.Contains(formNumberNames, "괄호:수") || slices.Contains(formNumberNames, "괄호:한글") {
		t.Fatal(formNumberNames)
	}
}
