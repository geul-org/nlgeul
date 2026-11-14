//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestWaysBracket — each way's inner bracket from its first line to the line of its last box's title, tagged with its rank and lambda, and the outer bracket tagged "어절경합" with the place of the stretch
package token

import "testing"

func TestWaysBracket(t *testing.T) {
	region := Span{0, 6}
	segs := []Segment{
		boxSegment(1, 2, -11.98, region, boxNode(0, 3, "아버지", "아버지"), boxNode(3, 6, "가방에", "가방에")),
		boxSegment(2, 2, -12.37, region, boxNode(0, 4, "아버지가", "아버지가"), boxNode(4, 6, "방에", "방에")),
	}
	ways := []boxWay{newBoxWay(segs[0], 0), newBoxWay(segs[1], 0)}
	lines := 0
	for i := range ways {
		ways[i].draw(40)
		lines += len(ways[i].lines)
	}
	b := waysBracket(ways, segs, lines)
	if lines != 12 || string(b.inner) != "┬││┘  ┬││┘  " || string(b.outer) != "┐││├││┘     " {
		t.Fatalf("%d %q %q", lines, string(b.inner), string(b.outer))
	}
	if b.tags[0] != "① λ 합 -11.98" || b.tags[6] != "② λ 합 -12.37" || b.tags[3] != "어절경합 [0-6]" {
		t.Fatalf("%q", b.tags)
	}
}
