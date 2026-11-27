//ff:func feature=parse type=reader control=sequence
//ff:what TestApproxAt — the approximate numbers that end at a place, each as its two numbers: one for each first number that ends where the second begins (십오 and 오 before 륙), the second as long as it runs (삼십 after 이), and none where the letters are no such pair
package parse

import (
	"reflect"
	"testing"
)

func TestApproxAt(t *testing.T) {
	if got := scanOf("칠팔년").approxAt(2); !reflect.DeepEqual(got, [][2]string{{"칠", "팔"}}) {
		t.Fatal(got)
	}
	if got := scanOf("이삼십명").approxAt(3); !reflect.DeepEqual(got, [][2]string{{"이", "삼십"}}) {
		t.Fatal(got)
	}
	if got := scanOf("십오륙세").approxAt(3); !reflect.DeepEqual(got, [][2]string{{"오", "륙"}, {"십오", "륙"}}) {
		t.Fatal(got)
	}
	if got := scanOf("삼십명").approxAt(2); got != nil {
		t.Fatal(got)
	}
	if got := scanOf("칠").approxAt(1); got != nil {
		t.Fatal(got)
	}
}
