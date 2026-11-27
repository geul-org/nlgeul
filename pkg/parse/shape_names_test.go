//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestShapeNames — the names the table of shapes gives a letter, one for each line (ⓐ 원:로마자, ？ 전각:기호); none for a letter that table lacks (○ is named by another table) and none for a letter in no table
package parse

import (
	"reflect"
	"testing"
)

func TestShapeNames(t *testing.T) {
	for r, want := range map[rune][]string{'ⓐ': {"원:로마자"}, '？': {"전각:기호"}, '○': nil, '가': nil} {
		if got := shapeNames(r); !reflect.DeepEqual(got, want) {
			t.Errorf("shapeNames(%c) = %q, want %q", r, got, want)
		}
	}
}
