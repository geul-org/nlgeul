//ff:func feature=parse type=helper control=sequence
//ff:what TestFormNames — a letter of the tables is named by its shape and what it writes, one name for each of its lines; a letter the tables lack has none
package parse

import (
	"slices"
	"testing"
)

func TestFormNames(t *testing.T) {
	if got := formNames('㈎'); !slices.Equal(got, []string{"괄호:한글"}) {
		t.Fatal(got)
	}
	if got := formNames('❶'); !slices.Equal(got, []string{"검은 원:수"}) {
		t.Fatal(got)
	}
	if got := formNames('가'); got != nil {
		t.Fatal(got)
	}
}
