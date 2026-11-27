//ff:func feature=data type=helper control=sequence
//ff:what run — makes the table of a mode (forms or groups) from the UnicodeData.txt and Blocks.txt at the two paths and writes it; stops when Blocks.txt is not of the version this program knows
package main

import (
	"bufio"
	"fmt"
	"io"
)

func run(mode, unicodeDataPath, blocksPath string, w io.Writer) error {
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
	write := map[string]func(io.Writer, *ucdTable, []blockRange) error{"forms": writeTable, "groups": writeGroups}[mode]
	if write == nil {
		return fmt.Errorf("명령을 모른다: %s (forms 또는 groups)", mode)
	}
	bw := bufio.NewWriter(w)
	if err := write(bw, t, blocks); err != nil {
		return err
	}
	return bw.Flush()
}
