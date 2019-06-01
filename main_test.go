package main

import (
	"bufio"
	"os"
	"testing"
	"bytes"
	"io/ioutil"
	"strings"
	"fmt"
)

func TestString(t *testing.T) {
	relPath()
	absPath()
}

func relPath() {
	var b bytes.Buffer
    w := bufio.NewWriter(&b)
	file, _ := os.Open("./fixtures/count.out")
	lcov(parseCoverage(bufio.NewReader(file), "", ""), w)
	t, _ := ioutil.ReadFile("./fixtures/lcov-rel.info")
	if strings.Compare(string(t), b.String()) != 0 {
		fmt.Print(b.String())
		panic("Relative path test failed.")
	}

}

func absPath() {
	var b bytes.Buffer
    w := bufio.NewWriter(&b)
	file, _ := os.Open("./fixtures/count.out")
	lcov(parseCoverage(bufio.NewReader(file), "/abs/path", ""), w)
	t, _ := ioutil.ReadFile("./fixtures/lcov-abs.info")
	if strings.Compare(string(t), b.String()) != 0 {
		fmt.Print(b.String())
		panic("Absolute path test failed.")
	}
}
