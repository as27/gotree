package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_isWhitespace(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{
			"tab",
			'	',
			true,
		},
		{
			"space",
			' ',
			true,
		},
		{
			"#",
			'#',
			true,
		},
		{
			"+",
			'+',
			true,
		},
		{
			"-",
			'-',
			true,
		},
		{
			"Letter",
			'a',
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWhitespace(tt.r); got != tt.want {
				t.Errorf("isWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLine(t *testing.T) {
	tests := []struct {
		name string
		l    string
		want lineElement
	}{
		{
			"no indent",
			"folder 1#",
			lineElement{
				indent: "",
				name:   "folder 1#",
			},
		},
		{
			"indent #",
			"# # #folder 1#",
			lineElement{
				indent: "# # #",
				name:   "folder 1#",
			},
		},
		{
			"indent tab",
			"		folder 1#",
			lineElement{
				indent: "		",
				name: "folder 1#",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLine(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

const simpleTreeIn2 = `folder
	folder1
		file 1
	folder2
		file 21
		folder21
			file 211
	file 3`

var simpleTree2Node = node{
	lineElement{"", "folder"},
	[]*node{
		&node{
			lineElement{"", "folder1"},
			[]*node{
				&node{
					lineElement{"", "file 1"},
					nil,
				},
			},
		},
		&node{
			lineElement{"", "folder2"},
			[]*node{
				&node{
					lineElement{"", "file 21"},
					[]*node{
						&node{
							lineElement{"", "folder21"},
							[]*node{
								&node{
									lineElement{"", "file 211"},
									nil,
								},
							},
						},
					},
				},
			},
		},
		&node{
			lineElement{"", "file 3"},
			nil,
		},
	},
}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name string
		r    io.Reader
		want node
	}{
		{
			"tree",
			strings.NewReader(simpleTreeIn2),
			simpleTree2Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.r); got.String() != tt.want.String() {
				t.Errorf("parseInput() = %s, want %s", got, tt.want)
			}
		})
	}
}
