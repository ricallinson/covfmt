package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"flag"
	"path/filepath"
)

type block struct {
	startLine  int
	startChar  int
	endLine    int
	endChar    int
	statements int
	covered    int
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

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	source := flag.String("f", "", "Path to the source coverage file.")
	output := flag.String("o", "", "Path for the output coverage file.")
	flag.Parse()
	if *source == "" {
		flag.PrintDefaults()
		return
	}
	if *output == "" {
		*output = filepath.Join(filepath.Dir(*source), "lcov.info")
	}
	coverage, err := os.Open(*source)
	if err != nil {
		log.Fatal("Error opening coverage file:", err)
	}
	defer coverage.Close()
	lcov(parseCoverage(coverage), *output)
}
