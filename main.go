package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var output io.Writer = os.Stdout

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("a txt file is expexted as input")
		os.Exit(2)
	}
	inFileName := args[0]
	inReader, err := os.Open(inFileName)
	if err != nil {
		fmt.Println("cannot open input file")
		os.Exit(1)
	}
	defer inReader.Close()
	err = convert(inReader, output)
	if err != nil {
		fmt.Printf("error converting: %s", err)
	}
}

func convert(r io.Reader, w io.Writer) error {
	return nil
}
