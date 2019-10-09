package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type node struct {
	element  lineElement
	children []*node
}

func (n node) String() string {
	out := fmt.Sprintf("%s\n", n.element.name)
	for _, n := range n.children {
		out = fmt.Sprintf("%s\n-->%s", out, *n)
	}
	return out
}

type lineElement struct {
	indent string
	name   string
}

func parseInput(r io.Reader) node {
	//var trees []node
	var path []*node
	var lastNode node
	var indentLevel int
	var root *node

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()
		le := parseLine(l)
		n := node{element: le}
		switch {
		case len(path) == 0:
			indentLevel = 0
			path = []*node{&n}
			root = &n
		case lastNode.element.indent == le.indent:
			parent := path[indentLevel-1]
			parent.children = append(parent.children, &n)
			path[indentLevel] = &n
		case strings.HasPrefix(le.indent, lastNode.element.indent):
			// one level deeper
			parent := path[indentLevel]
			parent.children = append(parent.children, &n)
			path[indentLevel].children = parent.children
			path = append(path, &n)
			indentLevel++
		default:
			// goes one or more levels up
			for i, p := range path {
				if p.element.indent == le.indent {
					path[i-1].children = append(path[i-1].children, &n)
					indentLevel = i
					path = path[:i]
					path = append(path, &n)
					break
				}
			}
		}
		fmt.Println(indentLevel, n.element.name)
		lastNode = n
	}
	return *root
}

func parseLine(l string) lineElement {
	var le lineElement
	for i, r := range l {
		if !isWhitespace(r) {
			le.name = l[i:]
			break
		}
		le.indent = le.indent + string(r)
	}
	return le
}

func isWhitespace(r rune) bool {
	const whitespces = " 	#-+"
	for _, w := range whitespces {
		if w == r {
			return true
		}
	}
	return false
}
