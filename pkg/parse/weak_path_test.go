//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestWeakPath — a unit written as a path with no head and one slash whose name holds Hangul and a sign that joins the parts of a name ends where a path ends (past the Hangul after a joining sign, else after the last letter that is no Hangul syllable); none when there is no slash, more than one, a head, no Latin letter at the front, Hangul before the slash, another sign, no Hangul in the name, or no joining sign between two letters of the name
package parse

import "testing"

func TestWeakPath(t *testing.T) {
	for in, want := range map[string]int{
		"notes/회의-결과": 11, "notes/memo-결과": 13, "docs/계약서_양식": 11, "notes/회의-v2를": 11, "notes/회의-결과를": 12,
		"notes": 0, "a/b/회의-결과": 0, "/회의-결과": 0, "회의/결과-값": 0, "3/회의-결과": 0, "notes/회의,결과-값": 0, "no,tes/회의-결과": 0,
		"notes/memo-v2": 0, "notes/memo-v2를": 0, "TV/라디오": 0, "notes/-회의": 0, "notes/회의-": 0, "notes/회의": 0, "": 0,
	} {
		if got := weakPath([]rune(in)); got != want {
			t.Errorf("%q: %d", in, got)
		}
	}
}
