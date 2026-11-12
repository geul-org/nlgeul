//ff:func feature=token type=helper control=sequence
//ff:what TestDistinct — each text once, in the order it first comes; nothing gives nothing
package token

import (
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	if got := distinct([]string{"b", "a", "b", "", "a", ""}); !reflect.DeepEqual(got, []string{"b", "a", ""}) {
		t.Fatal(got)
	}
	if got := distinct(nil); got != nil {
		t.Fatal(got)
	}
}
