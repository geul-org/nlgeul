package data

import (
	"embed"
	"io/fs"
)

// held holds the closed grammar tables of the folder tables, one
// tsv file a table. Only that folder is held: the word table
// (words.tsv) is a file given when the program is run.
//
//go:embed tables/*.tsv
var held embed.FS

// tables is where the tables are read from: the folder held in the
// program.
var tables fs.FS = held

// head begins the first line of a table, before the names of its
// columns; none is what a column with no value is written as, and joint what
// joins the values of a column that takes more than one.
const (
	head  = "# 칸(탭으로 나눔): "
	none  = "-"
	joint = "|"
)
