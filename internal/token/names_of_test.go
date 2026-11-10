//ff:func feature=token type=helper control=sequence
//ff:what TestNamesOf — joined names part into readings: two kinds with forms, the kinds of a word table's stem each under 열쇠, kinds alone with no form, a form that is itself a colon (cut at the first colon only), one word in two sets, and none for no names
package token

import (
	"reflect"
	"testing"
)

func TestNamesOf(t *testing.T) {
	if got := namesOf(""); got != nil {
		t.Fatalf("no names: %+v", got)
	}
	if got, want := namesOf("주격:가|보격:가"), []Name{{"주격", "가"}, {"보격", "가"}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
	if got, want := namesOf("열쇠:체언|부사"), []Name{{"열쇠", "체언"}, {"열쇠", "부사"}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
	if got, want := namesOf("수"), []Name{{"수", ""}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
	if got, want := namesOf("시각::|비율::"), []Name{{"시각", ":"}, {"비율", ":"}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
	if got, want := namesOf("관형사:두|보조 용언:두"), []Name{{"관형사", "두"}, {"보조 용언", "두"}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
	if got, want := namesOf("소수점|버전"), []Name{{"소수점", ""}, {"버전", ""}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
}
