package main

import (
	"bufio"
	"io"
	"strconv"
)

func writeLcovRecord(filePath string, blocks []*block, w *bufio.Writer) {

	w.WriteString("TN:\n")
	w.WriteString("SF:" + filePath + "\n")

	// Loop over functions
	// FN: line,name

	// FNF: total functions
	// FNH: covered functions

	// Loop over functions
	// FNDA: stats,name ?

	// Loop over lines
	total := 0
	covered := 0

	// Loop over each block and extract the lcov data needed.
	for _, b := range blocks {
		// For each line in a block we add an lcov entry and count the lines.
		for i := b.startLine; i <= b.endLine; i++ {
			total++
			if b.covered > 0 {
				covered++
			}
			w.WriteString("DA:" + strconv.Itoa(i) + "," + strconv.Itoa(b.covered) + "\n")
		}
	}

	w.WriteString("LF:" + strconv.Itoa(total) + "\n")
	w.WriteString("LH:" + strconv.Itoa(covered) + "\n")

	// Loop over branches
	// BRDA: ?

	// BRF: total branches
	// BRH: covered branches

	w.WriteString("end_of_record\n")
}

func lcov(blocks map[string][]*block, f io.Writer) {
	w := bufio.NewWriter(f)
	for file, fileBlocks := range blocks {
		writeLcovRecord(file, fileBlocks, w)
	}
	w.Flush()
}
