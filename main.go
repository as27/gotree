package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const version = "1.0.1"

var (
	flagImg     = flag.String("img", "", "Filename for an output image (png).")
	flagVersion = flag.Bool("v", false, "prints out the version")
)

func main() {
	flag.Parse()
	var output io.Writer = os.Stdout
	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}
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

	switch {
	case *flagImg != "":
		buf := &bytes.Buffer{}
		err = convert(inReader, buf)
		f, err := os.OpenFile(*flagImg, os.O_CREATE, 0666)
		defer f.Close()
		if err != nil {
			fmt.Println("cannot create img file: ", err)
			os.Exit(1)
		}
		drawText(f, readLines(buf), nil)
	default:
		err = convert(inReader, output)
		if err != nil {
			fmt.Printf("error converting: %s", err)
		}
	}

}

func convert(r io.Reader, w io.Writer) error {
	root := parseInput(r)
	walkTree(w, root, nil, true)
	return nil
}

func readLines(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func walkTree(w io.Writer, n node, indent []string, lastNode bool) {
	const (
		EMPTY  = "    "
		BRANCH = "├── "
		CORNER = "└── "
		LINE   = "│   "
	)
	if len(indent) > 0 {
		for i, ind := range indent[:len(indent)-1] {
			switch ind {
			case BRANCH:
				indent[i] = LINE
			case CORNER:
				indent[i] = EMPTY
			}
		}
	}
	fmt.Fprintf(w, "%s%s\n", strings.Join(indent, ""), n.element.name)
	last := false
	for i, nn := range n.children {
		branch := BRANCH
		if i == len(n.children)-1 {
			branch = CORNER
		}
		indentNew := append(indent, branch)
		walkTree(w, *nn, indentNew, last)
	}
}
