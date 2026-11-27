package main

// emojiBlocks are the blocks whose symbols are pictures.
var emojiBlocks = []string{"Miscellaneous Symbols and Pictographs", "Emoticons", "Transport and Map Symbols",
	"Supplemental Symbols and Pictographs", "Symbols and Pictographs Extended-A",
	"Mahjong Tiles", "Domino Tiles", "Playing Cards", "Chess Symbols", "Ornamental Dingbats"}

// mathBlocks are the blocks whose symbols are mathematical or technical.
var mathBlocks = []string{"Mathematical Operators", "Miscellaneous Mathematical Symbols-A",
	"Miscellaneous Mathematical Symbols-B", "Supplemental Mathematical Operators", "Miscellaneous Technical"}

// arrowBlocks are the blocks whose symbols are arrows.
var arrowBlocks = []string{"Arrows", "Supplemental Arrows-A", "Supplemental Arrows-B", "Supplemental Arrows-C",
	"Miscellaneous Symbols and Arrows"}

// shapeBlocks are the blocks whose symbols are shapes and bullets.
var shapeBlocks = []string{"Geometric Shapes", "Miscellaneous Symbols", "Dingbats", "Geometric Shapes Extended"}

// boxBlocks are the blocks whose symbols draw boxes and blocks.
var boxBlocks = []string{"Box Drawing", "Block Elements", "Symbols for Legacy Computing"}

// commonBlocks are the blocks of everyday punctuation and letters.
var commonBlocks = []string{"Basic Latin", "Latin-1 Supplement", "General Punctuation", "Supplemental Punctuation",
	"CJK Symbols and Punctuation", "Katakana", "Hiragana", "Spacing Modifier Letters"}

// groupCategories are the general categories whose characters the group table holds.
var groupCategories = []string{"Pc", "Pd", "Po", "Sm", "Sc", "Sk", "So"}

// directionPairs are two adjacent name words that together read as one direction.
var directionPairs = map[[2]string]string{{"LEFT", "RIGHT"}: "양쪽", {"UP", "DOWN"}: "위아래",
	{"NORTH", "EAST"}: "오른쪽 위", {"SOUTH", "EAST"}: "오른쪽 아래",
	{"NORTH", "WEST"}: "왼쪽 위", {"SOUTH", "WEST"}: "왼쪽 아래"}

// directionWords are the name words that read as one direction by themselves.
var directionWords = map[string]string{"RIGHTWARDS": "오른쪽", "RIGHT": "오른쪽", "LEFTWARDS": "왼쪽", "LEFT": "왼쪽",
	"UPWARDS": "위", "UP": "위", "DOWNWARDS": "아래", "DOWN": "아래"}
