//ff:func feature=parse type=helper control=sequence
//ff:what TestSeatsOf — the words of the best way of each region in turn make one line, each seat with the place of its region in the line; a region with no way gives no seat and the next region keeps its own place
package parse

import "testing"

func TestSeatsOf(t *testing.T) {
	line := sentenceRegions(t, "학생이 밥을 먹었다")
	line = append(line[:1], append([]*region{newRegion(Contest{Text: "가나"}, 3, nil)}, line[1:]...)...)
	seats := seatsOf(line)
	if len(seats) != 3 || seats[0].region != 0 || seats[1].region != 2 || seats[2].region != 3 {
		t.Fatal(seats)
	}
	if seats[0].near.Stem != "학생" || seats[1].near.Stem != "밥" || seats[2].near.Stem != "먹" || !seats[2].near.Pred {
		t.Fatal(seats[0].near, seats[1].near, seats[2].near)
	}
	if seatsOf(nil) != nil {
		t.Fatal("no region")
	}
}
