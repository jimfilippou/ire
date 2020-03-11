package utils

import (
	"encoding/json"
	"github.com/jimfilippou/ire/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func insertDocument(node *models.Node, index *int) error {
	log.Info(node.ID)
	return nil
}

func FeedTheDB() error {

	fileName := "/Users/jimfilippou/go/src/github.com/jimfilippou/ire/data/documents.json"
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
