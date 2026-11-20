//ff:func feature=parse type=helper control=sequence
//ff:what TestWrittenForms — the words and the forms a table writes them wrong in, in order, and not the forms written for other words; with the tables of this layer, 것, 거 and 꺼 (boundForms)
package parse

import (
	"slices"
	"testing"
)

func TestWrittenForms(t *testing.T) {
	if got := writtenForms([]string{"것", "거"}, map[string]string{"꺼": "거", "갯수": "개수"}); !slices.Equal(got, []string{"거", "것", "꺼"}) {
		t.Fatal(got)
	}
	if got := writtenForms([]string{"거"}, nil); !slices.Equal(got, []string{"거"}) {
		t.Fatal(got)
	}
	if !slices.Equal(boundForms, []string{"거", "것", "꺼"}) {
		t.Fatal(boundForms)
	}
}
