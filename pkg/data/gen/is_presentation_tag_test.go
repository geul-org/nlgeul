//ff:func feature=data type=helper control=sequence
//ff:what TestIsPresentationTag — the four Arabic presentation tags stand, others do not
package main

import "testing"

func TestIsPresentationTag(t *testing.T) {
	if !isPresentationTag("<isolated>") || isPresentationTag("<circle>") {
		t.Fatal("tags")
	}
}
