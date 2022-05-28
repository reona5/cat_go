package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	isNumber bool
)

func init() {
	flag.BoolVar(&isNumber, "n", false, "Line number is visible when it's true.")
}

func main() {
	flag.Parse()

	filenames := flag.Args()
	for _, fn := range filenames {
		readFile(fn)
	}
}

func readFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ファイルの読み込みに失敗しました", err)
		os.Exit(1)
	}
	scanLines(f)
	defer f.Close()
}

func scanLines(f *os.File) {
	line := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if isNumber {
			fmt.Fprintln(os.Stdout, line, scanner.Text())
			line++
		} else {
			fmt.Fprintln(os.Stdout, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ファイルの読み込みに失敗しました", err)
		os.Exit(1)
	}
}
