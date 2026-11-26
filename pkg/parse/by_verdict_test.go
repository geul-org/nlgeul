//ff:func feature=parse type=helper control=sequence
//ff:what TestByVerdict — the readings of a sign of a pair in the order of their verdicts, the others as they came: a lowered reading goes behind with its way, a reading as the letter alone counts as 1 and gives a way but no candidate, a kind and depth met twice is kept once with its first (highest) verdict, each way once; nothing gives nothing
package parse

import (
	"reflect"
	"testing"
)

func TestByVerdict(t *testing.T) {
	reads := []pairWay{
		{pairing{openName, 1, 0.25}, "\""},
		{pairing{shutName, 0, 1}, ""},
		{pairing{verdict: 1}, "("},
		{pairing{shutName, 0, 0.5}, "["},
		{pairing{openName, 1, 1}, "\""},
		{pairing{openName, 2, 0.25}, "(\""},
	}
	candidates, ways := byVerdict(reads)
	if want := []pairing{{shutName, 0, 1}, {openName, 1, 1}, {openName, 2, 0.25}}; !reflect.DeepEqual(candidates, want) {
		t.Fatal(candidates)
	}
	if want := []string{"", "(", "\"", "[", "(\""}; !reflect.DeepEqual(ways, want) {
		t.Fatal(ways)
	}
	if reads[0].read.verdict != 0.25 || reads[1].way != "" {
		t.Fatal("the readings given are not put in another order", reads)
	}
	if candidates, ways = byVerdict(nil); candidates != nil || ways != nil {
		t.Fatal(candidates, ways)
	}
	if candidates, ways = byVerdict([]pairWay{{pairing{verdict: 1}, "'"}}); candidates != nil || !reflect.DeepEqual(ways, []string{"'"}) {
		t.Fatal("the letter alone", candidates, ways)
	}
}
