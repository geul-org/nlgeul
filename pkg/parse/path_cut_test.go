//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestPathCut — a path that ends in a slash with Hangul alone after it up to the unit's end is cut at its end; no cut when the unit begins with no path, the path is a slash alone, it does not end in a slash, nothing follows it, or something that is no Hangul syllable follows
package parse

import "testing"

func TestPathCut(t *testing.T) {
	cases := []struct {
		in, name string
		end, cut int
	}{
		{"~/a/b/회의록", pathName, 6, 6}, {"~/a/b/는", pathName, 6, 6}, {"../a/폴더는", pathName, 5, 5},
		{"~/a/b/회의록", "주소", 6, 0}, {"/좌측", pathName, 1, 0}, {"~/a/b를", pathName, 5, 0}, {"~/a/b/", pathName, 6, 0}, {"~/a/b/회의록", pathName, 0, 0}, {"~/a/b/회의-록", pathName, 6, 0}, {"~/a/b/회의,", pathName, 6, 0},
	}
	for _, c := range cases {
		if got := pathCut([]rune(c.in), c.name, c.end); got != c.cut {
			t.Errorf("%s %s %d: %d", c.in, c.name, c.end, got)
		}
	}
}
