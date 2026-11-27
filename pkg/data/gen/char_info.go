//ff:type feature=data type=model
//ff:what charInfo — the columns of one UnicodeData.txt line that the table needs: name, general category, decomposition, numeric value
package main

type charInfo struct {
	Name     string
	Category string
	Decomp   string
	Numeric  string
}
