//ff:type feature=data type=model
//ff:what Row — one line of a table: the file it is in, its number in the file, the names of the columns as the file's first line gives them, and its columns
package data

type Row struct {
	File  string
	Line  int
	Names []string
	Cells []string
}
