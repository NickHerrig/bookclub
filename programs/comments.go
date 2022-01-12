package main

import (
	_ "embed"
	"fmt"
	"regexp"
)

//go:embed comments.go
var self string

var (
	lineCommentPattern  = regexp.MustCompile(`(?m)//.*$`)
	blockCommentPattern = regexp.MustCompile(`(?s)/\*(.*?)\*/`)
)

// this is a line comment

/* this is a block comment */

/*
this
is is big
block
comment
*/

func main() {
	// easy quine
//	fmt.Println(self)

	lines := lineCommentPattern.FindAllString(self, -1)
	for _, match := range lines {
		fmt.Println(match)
	}

	blocks := blockCommentPattern.FindAllString(self, -1)
	for _, match := range blocks {
		fmt.Println(match)
	}
}
