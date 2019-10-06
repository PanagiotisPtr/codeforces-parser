package main

import (
	"fmt"
	"os"

	"github.com/panagiotisptr/codeforces-parser/pkg/builder"
)

func main() {
	competitionUrl := os.Args[1]
	fmt.Println("Parsing competition")
	builder.BuildCompetition(competitionUrl)
	fmt.Println("Done")
}