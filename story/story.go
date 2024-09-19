package story

import (
	"encoding/json"
	"fmt"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"string"`
}

func ParseStoryJson(filepath string) (Story, error) {
	var result Story

	file, err := os.Open(filepath)
	if err != nil {
		return result, fmt.Errorf("failed to open the file %s:\n\t%v", filepath, err)
	}
	defer file.Close()

	data := json.NewDecoder(file)

	if errU := data.Decode(&result); errU != nil {
		return result, fmt.Errorf("failed to unmarshal json of the file %s:\n\t%v", filepath, errU)
	}

	return result, nil
}
