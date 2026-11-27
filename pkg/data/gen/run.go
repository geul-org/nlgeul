//ff:func feature=data type=helper control=sequence
//ff:what run — makes the special-character table from the UnicodeData.txt and Blocks.txt at the two paths and writes it; stops when Blocks.txt is not of the version this program knows
package main

import (
	"bufio"
	"fmt"
	"io"
)

func run(unicodeDataPath, blocksPath string, w io.Writer) error {
	t, err := loadUCD(unicodeDataPath)
	if err != nil {
		return err
	}
	blocks, version, err := loadBlocks(blocksPath)
	if err != nil {
		return err
	}
	if version != unicodeVersion {
		return fmt.Errorf("유니코드 판이 다르다: 프로그램 %s, Blocks.txt %s", unicodeVersion, version)
	}
	bw := bufio.NewWriter(w)
	if err := writeTable(bw, t, blocks); err != nil {
		return err
	}
	return bw.Flush()
}
