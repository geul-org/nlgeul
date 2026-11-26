//ff:func feature=parse type=helper control=sequence
//ff:what TestEndsNamed — a letter that ends one of the closings named is one a particle of theirs is looked for after (다 of 맺음:다), a letter that ends none is not, any letter is when one of them is fused into the syllable before it, and no list takes none
package parse

import "testing"

func TestEndsNamed(t *testing.T) {
	names := []string{"맺음:다", "맺음:라"}
	if !endsNamed(names, '다') || !endsNamed(names, '라') {
		t.Fatal("the last letter of a closing named")
	}
	if endsNamed(names, '고') || endsNamed(nil, '다') {
		t.Fatal("a letter that ends none of them")
	}
	if !endsNamed([]string{"맺음:게", "맺음:어(합침)"}, '가') {
		t.Fatal("a closing fused into the syllable before it")
	}
}
