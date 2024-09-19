package main

import (
	"flag"
	"fmt"
	"log"

	"cyoa/story"
)

func main() {
	// using flag for file
	storyFile := flag.String("f", "../../gopher.json", "the JSON file with CYOA story")
	flag.Parse()

	storyParsed, err := story.ParseStoryJson(*storyFile)
	if err != nil {
		log.Fatalf("failed to parse the story file %s:\n\t%v", *storyFile, err)
	}

	fmt.Printf("%+v\n", storyParsed)
}
