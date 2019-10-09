package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

const simpleTreeIn = `folder
	folder1
	folder2
		file 21
		file 22
	folder3
		file 31`

const simpleTree = `folder
├── folder1
├── folder2
│   ├── file 21
│   └── file 22
└── folder3
    └── file 31
`

func Test_convert(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		r       io.Reader
		wantW   string
		wantErr bool
	}{
		{
			"simple tree",
			strings.NewReader(simpleTreeIn),
			simpleTree,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := convert(tt.r, w); (err != nil) != tt.wantErr {
				t.Errorf("convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("convert() = \n%v, want \n%v,", gotW, tt.wantW)
			}
		})
	}
}
