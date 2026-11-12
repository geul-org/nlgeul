//ff:func feature=token type=helper control=sequence
//ff:what TestMarkNames — the names of some candidates of a sign come back in order; no candidates, no names
package token

import (
	"reflect"
	"testing"
)

func TestMarkNames(t *testing.T) {
	got := markNames([]markRead{{"파일명", 1}, {"도메인", 0.5}})
	if !reflect.DeepEqual(got, []string{"파일명", "도메인"}) {
		t.Fatal(got)
	}
	if got := markNames(nil); len(got) != 0 {
		t.Fatal(got)
	}
}
