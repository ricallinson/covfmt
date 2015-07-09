package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type block struct {
	startLine  int
	startChar  int
	endLine    int
	endChar    int
	statements int
	covered    int
}

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

func parseCoverageLine(line string) (string, *block, bool) {
	if line == "mode: set" {
		return "", nil, false
	}
	path := strings.Split(line, ":")
	parts := strings.Split(path[1], " ")
	sections := strings.Split(parts[0], ",")
	start := strings.Split(sections[0], ".")
	end := strings.Split(sections[1], ".")
	// Populate the block.
	b := &block{}
	b.startLine, _ = strconv.Atoi(start[0])
	b.startChar, _ = strconv.Atoi(start[1])
	b.endLine, _ = strconv.Atoi(end[0])
	b.endChar, _ = strconv.Atoi(end[1])
	b.statements, _ = strconv.Atoi(parts[1])
	b.covered, _ = strconv.Atoi(parts[2])
	// Remove the underscore (_) from the beginning of the path.
	return path[0][1:], b, true
}

func parseCoverage(coverage io.Reader) map[string][]*block {
	scanner := bufio.NewScanner(coverage)
	blocks := map[string][]*block{}
	for scanner.Scan() {
		if f, b, ok := parseCoverageLine(scanner.Text()); ok {
			// Make sure the filePath is a key in the map.
			if _, ok := blocks[f]; ok == false {
				blocks[f] = []*block{}
			}
			blocks[f] = append(blocks[f], b)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
	return blocks
}

func main() {
	lcov(parseCoverage(os.Stdin), os.Stdout)
}
