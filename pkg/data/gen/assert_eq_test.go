//ff:func feature=data type=helper control=sequence
//ff:what assertEq — fails the test when two strings differ
package main

import "testing"

func assertEq(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
