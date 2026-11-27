//ff:func feature=data type=helper control=iteration dimension=1
//ff:what TestArrowDirection — pairs, single words and diagonals each read as one direction, and no direction or two different directions give empty
package main

import "testing"

func TestArrowDirection(t *testing.T) {
	cases := map[string]string{
		"RIGHTWARDS ARROW":                      "오른쪽",
		"LEFTWARDS DOUBLE ARROW":                "왼쪽",
		"UPWARDS ARROW":                         "위",
		"DOWN ARROW":                            "아래",
		"LEFT RIGHT ARROW":                      "양쪽",
		"UP DOWN ARROW":                         "위아래",
		"SOUTH EAST ARROW":                      "오른쪽 아래",
		"NORTH EAST ARROW":                      "오른쪽 위",
		"NORTH WEST ARROW":                      "왼쪽 위",
		"SOUTH WEST ARROW":                      "왼쪽 아래",
		"LEFTWARDS ARROW OVER RIGHTWARDS ARROW": "",
		"CIRCLED ARROW":                         "",
		"RIGHTWARDS ARROW TO BAR WITH RIGHTWARDS TAIL": "오른쪽",
	}
	for name, want := range cases {
		assertEq(t, arrowDirection(name), want)
	}
}
