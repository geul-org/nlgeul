//ff:func feature=data type=helper control=selection
//ff:what skippedCategory — whether a general category is one the table never holds: unassigned, private use, surrogate or control
package main

func skippedCategory(cat string) bool {
	switch cat {
	case "Cn", "Co", "Cs", "Cc":
		return true
	}
	return false
}
