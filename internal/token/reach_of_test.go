//ff:func feature=token type=helper control=sequence
//ff:what TestReachOf — the root or a seed is worth 0, a state with a value has it, a state with none yet has none
package token

import "testing"

func TestReachOf(t *testing.T) {
	rc := &reach{value: map[*path]float64{}}
	root, a, b := &path{}, &path{prev: &path{}}, &path{prev: &path{}}
	rc.value[a] = -2.5
	if v, ok := rc.of(root); !ok || v != 0 {
		t.Fatal("a root")
	}
	if v, ok := rc.of(a); !ok || v != -2.5 {
		t.Fatal("a value")
	}
	if _, ok := rc.of(b); ok {
		t.Fatal("none yet")
	}
}
