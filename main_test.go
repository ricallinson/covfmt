package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestRelPath(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	file, _ := os.Open("./fixtures/count.out")
	lcov(parseCoverage(bufio.NewReader(file), "", ""), w)
	test, _ := ioutil.ReadFile("./fixtures/lcov-rel.info")
	if bytes.Compare(test, b.Bytes()) != 0 {
		t.Errorf("Relative path test failed.")
	}

}

func TestAbsPath(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	file, _ := os.Open("./fixtures/count.out")
	lcov(parseCoverage(bufio.NewReader(file), "/abs/path", ""), w)
	test, _ := ioutil.ReadFile("./fixtures/lcov-abs.info")
	if bytes.Compare(test, b.Bytes()) != 0 {
		t.Errorf("Absolute path test failed.")
	}
}

func TestTrimPath(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	file, _ := os.Open("./fixtures/count.out")
	lcov(parseCoverage(bufio.NewReader(file), "", "github.com/ricallinson/mk3"), w)
	test, _ := ioutil.ReadFile("./fixtures/lcov-trim.info")
	if bytes.Compare(test, b.Bytes()) != 0 {
		t.Errorf("Absolute path test failed.")
	}
}

func TestAbsTrimPath(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	file, _ := os.Open("./fixtures/count.out")
	lcov(parseCoverage(bufio.NewReader(file), "/abs/path", "github.com/ricallinson/mk3"), w)
	test, _ := ioutil.ReadFile("./fixtures/lcov-abs-trim.info")
	if bytes.Compare(test, b.Bytes()) != 0 {
		t.Errorf("Absolute path test failed.")
	}
}
