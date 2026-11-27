//ff:func feature=data type=reader control=sequence
//ff:what loadUCD — reads the UnicodeData.txt file at a path into a table
package main

import "os"

func loadUCD(path string) (*ucdTable, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return parseUnicodeData(f)
}
