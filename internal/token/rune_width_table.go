package token

// zeroCells are the letters that take no cell of a terminal (combining
// marks, zero width marks, variation selectors); wideCells are the East
// Asian wide letters, which take two. Every other letter takes one,
// the box drawing letters too. The ranges are the ones the box output
// is drawn with, not the whole of Unicode.
var zeroCells = []runeRange{
	{0x0300, 0x036F}, {0x200B, 0x200F}, {0xFE00, 0xFE0F},
}

var wideCells = []runeRange{
	{0x1100, 0x115F}, {0x2E80, 0x303E}, {0x3041, 0x33FF}, {0x3400, 0x4DBF},
	{0x4E00, 0x9FFF}, {0xA000, 0xA4CF}, {0xAC00, 0xD7A3}, {0xF900, 0xFAFF},
	{0xFE30, 0xFE4F}, {0xFF00, 0xFF60}, {0xFFE0, 0xFFE6}, {0x1F300, 0x1F64F},
	{0x1F900, 0x1F9FF}, {0x20000, 0x3FFFD},
}
