//ff:func feature=data type=helper control=iteration dimension=1
//ff:what TestGroupOf — each rule gives its group and name, an arrow with a readable direction is 화살표, and an arrow without one stays 기호 화살표
package main

import "testing"

func TestGroupOf(t *testing.T) {
	tbl := &ucdTable{Chars: map[rune]charInfo{
		0x24:    {Name: "DOLLAR SIGN", Category: "Sc"},
		0x1F600: {Name: "GRINNING FACE", Category: "So"},
		0x2500:  {Name: "BOX DRAWINGS LIGHT HORIZONTAL", Category: "So"},
		0x2192:  {Name: "RIGHTWARDS ARROW", Category: "Sm"},
		0x21C6:  {Name: "LEFTWARDS ARROW OVER RIGHTWARDS ARROW", Category: "So"},
		0x2B00:  {Name: "WHITE SQUARE", Category: "So"},
		0x2200:  {Name: "FOR ALL", Category: "Sm"},
		0x2B:    {Name: "PLUS SIGN", Category: "Sm"},
		0x25A0:  {Name: "BLACK SQUARE", Category: "So"},
		0x26:    {Name: "AMPERSAND", Category: "Po"},
	}}
	cases := []struct {
		r           rune
		block       string
		group, name string
	}{
		{0x24, "Basic Latin", "기호", "돈"},
		{0x1F600, "Emoticons", "기호", "이모지"},
		{0x2500, "Box Drawing", "기호", "상자 그림"},
		{0x2192, "Arrows", "화살표", "오른쪽"},
		{0x21C6, "Arrows", "기호", "화살표"},
		{0x2B00, "Arrows", "기호", "화살표"},
		{0x2200, "Mathematical Operators", "기호", "수학"},
		{0x2B, "Basic Latin", "기호", "수학"},
		{0x25A0, "Geometric Shapes", "기호", "도형"},
		{0x26, "Basic Latin", "기호", "그 밖"},
	}
	for _, c := range cases {
		g, n := groupOf(tbl, c.r, c.block)
		assertEq(t, g+"/"+n, c.group+"/"+c.name)
	}
	tbl.Chars[0x2B00] = charInfo{Name: "RIGHTWARDS ALPHA", Category: "So"}
	if g, n := groupOf(tbl, 0x2B00, "Arrows"); g != "기호" || n != "화살표" {
		t.Fatal(g, n)
	}
}
