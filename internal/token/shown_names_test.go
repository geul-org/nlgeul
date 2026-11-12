//ff:func feature=token type=helper control=sequence
//ff:what TestShownNames — candidates are shown without their sign, a thousands separator as a number, and none twice (the higher verdict stays); each keeps its own verdict
package token

import (
	"reflect"
	"testing"
)

func TestShownNames(t *testing.T) {
	if got := shownNames([]markRead{{"시각::", 1}, {"비율::", 0.5}}); !reflect.DeepEqual(got, []markRead{{"시각", 1}, {"비율", 0.5}}) {
		t.Fatal(got)
	}
	if got := shownNames([]markRead{{"자릿수:,", 1}}); !reflect.DeepEqual(got, []markRead{{numName, 1}}) {
		t.Fatal(got)
	}
	if got := shownNames([]markRead{{"버전:.", 0.4}, {"IP:.", 1}, {"버전:/", 0.9}}); !reflect.DeepEqual(got, []markRead{{"버전", 0.9}, {"IP", 1}}) {
		t.Fatal(got)
	}
}
