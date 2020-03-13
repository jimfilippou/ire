/*
 * Copyright (c) 2020.
 * Jim Filippou · jimfilippou8@gmail.com
 */

package utils

import (
	"encoding/json"
	"github.com/jimfilippou/ire/models"
	"io/ioutil"
	"os"
	"strings"
)

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// Gets a text file as an input, splits the contents and assigns them
// to nodes, which will be later iterated through
func getNodes(data []byte) []models.Node {
	var nodes []models.Node
	text := string(data)
	chunks := strings.Split(text, " /// ")

	// Read the first node
	nodes = append(nodes, models.Node{
		ID:   strings.Split(chunks[0], "\n")[0],
		Text: strings.Split(chunks[0], "\n")[1],
	})

	for i := 1; i <= len(chunks)-1; i++ {
		tokens := strings.Split(chunks[i], "\n")
		nodes = append(nodes, models.Node{
			ID:   tokens[1],
			Text: tokens[2],
		})
	}
	return nodes
}

//	This public function converts "documents.txt" to "documents.json"
//	in order to make them accessible to the "Feed" utility.
func CreateFile() error {

	fileName := "/Users/jimfilippou/go/src/github.com/jimfilippou/ire/data/documents.txt"
	outputFi := "/Users/jimfilippou/go/src/github.com/jimfilippou/ire/data/documents.json"

	err := checkFile(fileName)
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	nodes := getNodes(file)

	output, err := json.MarshalIndent(nodes, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFi, output, 0644)
	if err != nil {
		return err
	}

	return nil
}
