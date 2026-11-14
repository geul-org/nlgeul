//ff:func feature=token type=reader control=sequence
//ff:what TestUnion — names of both lists once each, in order; the first list is not changed
package token

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	first := []string{"도메인"}
	if got := union(nil, []string{"파일명", "도메인"}); !reflect.DeepEqual(got, []string{"파일명", "도메인"}) {
		t.Fatal(got)
	}
	if got := union(first, []string{"파일명", "도메인"}); !reflect.DeepEqual(got, []string{"도메인", "파일명"}) || len(first) != 1 {
		t.Fatal(got, first)
	}
	if got := union([]string{"버전"}, []string{"버전"}); !reflect.DeepEqual(got, []string{"버전"}) {
		t.Fatal(got)
	}
}
