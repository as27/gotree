package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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
	root := parseInput(r)
	walkTree(w, root, nil, true)
	return nil
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
