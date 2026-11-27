//ff:func feature=data type=helper control=sequence
//ff:what TestSkippedCategory — the four C categories are skipped, letters and symbols are not
package main

import "testing"

func TestSkippedCategory(t *testing.T) {
	if !skippedCategory("Cn") || !skippedCategory("Cc") || skippedCategory("So") {
		t.Fatal("categories")
	}
}
