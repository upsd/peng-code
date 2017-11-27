package main

import (
	"os"
	"fmt"
	"./match-finder"
)

func main() {
	slang := make(map[string]string)
	slang["mash man"] = "gunman"
	slang["arms"] = "strong"
	slang["beast"] = "cool"
	slang["bare"] = "a lot"
	slang["tourist"] = "clueless person"
	slang["peng"] = "good looking"

	terms := os.Args[1:]

	matches := match_finder.CheckForMatches(terms, slang)

	for k, v := range matches {
		fmt.Printf("%s -> %s \n", k, v)
	}
}