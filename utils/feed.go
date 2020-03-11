package utils

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io/ioutil"
	"ire/models"
	"strings"
)

func insertDocument(node *models.Node, index *int) error {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}

	_, _ = es.Create("ire", node.ID, strings.NewReader(fmt.Sprintf("%v", node)))

	return nil
}

func FeedTheDB() error {

	fileName := "C:\\Users\\Elias\\go\\src\\ire\\data\\documents.json"
	var nodes []models.Node

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &nodes)
	if err != nil {
		return err
	}

	for index, node := range nodes {
		err := insertDocument(&node, &index)
		if err != nil {
			return err
		}
	}

	return nil
}
