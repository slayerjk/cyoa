package main

import (
	"cyoa/internal/story"
	"flag"
	"fmt"
	"html/template"
	"log"
)

const htmlTemplatePath = "template.html"

func main() {
	// using flag for file
	storyFile := flag.String("f", "gopher.json", "the JSON file with CYOA story")
	serverPort := flag.Int("p", 3000, "the port to start the server on")
	flag.Parse()

	storyParsed, err := story.ParseStoryJson(*storyFile)
	if err != nil {
		log.Fatalf("failed to parse the story file %s:\n\t%v", *storyFile, err)
	}

	fmt.Printf("%+v\n", storyParsed)

	templ, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("failed to parse the template file %s:\n\t%v", htmlTemplatePath, err)
	}
	fmt.Println(templ)
}
