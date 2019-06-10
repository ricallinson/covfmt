package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type block struct {
	filepath   string
	startLine  int
	startChar  int
	endLine    int
	endChar    int
	statements int
	covered    int
}

func writeLcovRecord(blocks []*block, w *bufio.Writer) {

	w.WriteString("TN:\n")
	w.WriteString("SF:" + blocks[0].filepath + "\n")

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

func lcov(blocks []*block, f io.Writer) {
	var start int
	cur := blocks[0].filepath
	w := bufio.NewWriter(f)
	for i, b := range blocks {
		if strings.Compare(b.filepath, cur) != 0 {
			writeLcovRecord(blocks[start:i-1], w)
			start = i
			cur = b.filepath
		}
	}
	w.Flush()
}

func parseCoverageLine(line string, prefix string, cutset string) (*block, bool) {
	if strings.HasPrefix(line, "mode:") {
		return nil, false
	}
	path := strings.Split(line, ":")
	parts := strings.Split(path[1], " ")
	sections := strings.Split(parts[0], ",")
	start := strings.Split(sections[0], ".")
	end := strings.Split(sections[1], ".")
	// Populate the block.
	b := &block{}
	b.filepath = filepath.Join(prefix, path[0])
	b.startLine, _ = strconv.Atoi(start[0])
	b.startChar, _ = strconv.Atoi(start[1])
	b.endLine, _ = strconv.Atoi(end[0])
	b.endChar, _ = strconv.Atoi(end[1])
	b.statements, _ = strconv.Atoi(parts[1])
	b.covered, _ = strconv.Atoi(parts[2])
	// Remove the "cutset" string from the beginning of the path if the CLI option is present.
	if len(cutset) > 0 {
		path[0] = strings.TrimLeft(path[0], cutset)
	}
	return b, true
}

func parseCoverage(coverage io.Reader, prefix string, cutset string) []*block {
	scanner := bufio.NewScanner(coverage)
	blocks := []*block{}
	for scanner.Scan() {
		if b, ok := parseCoverageLine(scanner.Text(), prefix, cutset); ok {
			blocks = append(blocks, b)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
	return blocks
}

func main() {
	var cutset string
	flag.StringVar(&cutset, "trim", "", "An optional string that will be trimmed from the front of the source file name.")
	flag.Parse()
	var prefix string
	if len(os.Args) == 2 {
		prefix = os.Args[1]
	}
	lcov(parseCoverage(os.Stdin, prefix, cutset), os.Stdout)
}
