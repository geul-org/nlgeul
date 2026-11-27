package main

import "regexp"

// The characters come from the Unicode Character Database 15.0.0
// (Unicode License v3, https://www.unicode.org/license.txt). The files
// UnicodeData.txt and Blocks.txt are taken from
// https://www.unicode.org/Public/15.0.0/ucd/ and are not kept in this repository.
const unicodeVersion = "15.0.0"

// formsKept are the forms the table holds; the other forms (compatibility
// jamo, typefaces, half-width, vertical, small, tortoise-shell brackets,
// other compatibility, no-break) are not held. To hold more, add the form here
// and make the table again.
var formsKept = []string{"원", "괄호", "검은 원", "겹원", "네모", "검은 네모", "점 붙음", "쉼표 붙음",
	"로마 숫자", "분수", "위첨자", "아래첨자", "전각"}

// presentationTags are the decomposition tags of the Arabic presentation forms; they are not special.
var presentationTags = []string{"<initial>", "<medial>", "<final>", "<isolated>"}

var enclosedBlocks = []string{"Enclosed Alphanumerics", "Enclosed CJK Letters and Months",
	"Enclosed Alphanumeric Supplement", "Enclosed Ideographic Supplement"}

// numberBlocks are the blocks whose number characters with no decomposition are held as numbers.
var numberBlocks = []string{"Number Forms", "Latin-1 Supplement", "Superscripts and Subscripts", "Dingbats"}

// nameForms are the forms read from the name of a character, longest key first.
var nameForms = [][2]string{{"NEGATIVE CIRCLED", "검은 원"}, {"DOUBLE CIRCLED", "겹원"},
	{"NEGATIVE SQUARED", "검은 네모"}, {"CIRCLED", "원"}, {"PARENTHESIZED", "괄호"}, {"SQUARED", "네모"}}

var tagForms = map[string]string{"<circle>": "원", "<wide>": "전각", "<narrow>": "반각", "<super>": "위첨자",
	"<sub>": "아래첨자", "<fraction>": "분수", "<square>": "네모", "<small>": "작은 꼴", "<vertical>": "세로 꼴",
	"<font>": "글자체", "<noBreak>": "줄 안 바꿈"}

var blockLine = regexp.MustCompile(`^([0-9A-F]+)\.\.([0-9A-F]+); (.+)`)
var versionLine = regexp.MustCompile(`^# Blocks-([0-9.]+)\.txt`)
var latinLetterName = regexp.MustCompile(`LATIN (CAPITAL|SMALL) LETTER ([A-Z])$`)
