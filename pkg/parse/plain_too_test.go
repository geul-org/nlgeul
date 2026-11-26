//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestPlainToo — two names joined by one slash and read as a path for it keep the new word of the same letters after the path (I/O, src/main, .cache/db); a path with a head or a second slash, an atom with another name or two names, and no candidate stay as they are
package parse

import "testing"

func TestPlainToo(t *testing.T) {
	path := []stemCandidate{{kindCandidate{0, pathName, 1}, "I/O", ""}}
	if got := plainToo("I/O", path); len(got) != 2 || got[0].name != pathName || got[1].name != "" || got[1].v != 0 || got[1].stem != "I/O" {
		t.Fatal(got)
	}
	if got := plainToo(".cache/db", path[:1]); len(got) != 2 {
		t.Fatal(got)
	}
	for _, stem := range []string{"/usr", "~/a", "a/b/c", "ab"} {
		if got := plainToo(stem, path[:1]); len(got) != 1 {
			t.Error(stem, got)
		}
	}
	if got := plainToo("a/b", []stemCandidate{{kindCandidate{0, "붙임", 1}, "a/b", ""}}); len(got) != 1 {
		t.Fatal(got)
	}
	two := []stemCandidate{{kindCandidate{0, pathName, 1}, "a/b.md", ""}, {kindCandidate{0, "파일명", 1}, "a/b.md", ""}}
	if got := plainToo("a/b.md", two); len(got) != 2 {
		t.Fatal(got)
	}
	if got := plainToo("a/b", nil); got != nil {
		t.Fatal(got)
	}
}
