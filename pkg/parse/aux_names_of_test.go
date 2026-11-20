//ff:func feature=parse type=helper control=sequence
//ff:what TestAuxNamesOf — each stem is named as the auxiliary it is read under (보조 용언:있), in the order of the stems; none gives an empty list
package parse

import (
	"slices"
	"testing"
)

func TestAuxNamesOf(t *testing.T) {
	if got := auxNamesOf([]string{"있", "보", "주"}); !slices.Equal(got, []string{"보조 용언:있", "보조 용언:보", "보조 용언:주"}) {
		t.Fatal(got)
	}
	if got := auxNamesOf(nil); got == nil || len(got) != 0 {
		t.Fatal(got)
	}
}
