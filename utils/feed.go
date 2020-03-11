package utils

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io/ioutil"
	"ire/models"
	"log"
	"strconv"
	"strings"
)

func insertDocument(node *models.Node, index *int) error {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      "ire",
		DocumentID: strconv.Itoa(*index + 1),
		Body:       strings.NewReader("WOAG!"),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.New(res.Status() + " Error indexing document ID=" + node.ID)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {

			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

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
