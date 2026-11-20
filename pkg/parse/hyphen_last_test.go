//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestHyphenLast — 붙임 goes after the other names (파일명, 도메인), which keep their order; names with no 붙임, 붙임 alone and none stay as they are
package parse

import (
	"reflect"
	"testing"
)

func TestHyphenLast(t *testing.T) {
	cases := []struct{ in, want []string }{
		{[]string{"붙임", "파일명"}, []string{"파일명", "붙임"}},
		{[]string{"붙임", "파일명", "도메인"}, []string{"파일명", "도메인", "붙임"}},
		{[]string{"경로", "붙임", "파일명"}, []string{"경로", "파일명", "붙임"}},
		{[]string{"파일명", "도메인"}, []string{"파일명", "도메인"}},
		{[]string{"붙임"}, []string{"붙임"}},
		{nil, nil},
	}
	for _, c := range cases {
		if got := hyphenLast(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%v: %v, want %v", c.in, got, c.want)
		}
	}
}
