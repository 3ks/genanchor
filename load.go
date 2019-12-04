package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Load() (ab map[string]string, sp map[string]bool) {
	return loadAbbreviation(), loadSpelling()
}

func loadAbbreviation() (ab map[string]string) {
	ab = make(map[string]string)
	f, err := os.Open(".abbreviation")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		v, _, e := rd.ReadLine()
		if e == io.EOF {
			break
		}
		values := strings.SplitAfterN(string(v), "=", 2)
		ab[values[0]] = values[1]
	}
	return
}

func loadSpelling() (sp map[string]bool) {
	sp = make(map[string]bool)
	f, err := os.Open(".spelling")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		v, _, e := rd.ReadLine()
		if e == io.EOF {
			break
		}
		sp[string(v)] = true
	}
	return
}
