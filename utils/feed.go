/*
 * Copyright (c) 2020.
 * Jim Filippou · jimfilippou8@gmail.com
 */

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jimfilippou/ire/models"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"time"
)

func FeedTheDB() error {

	fileName := filepath.Join("utils/../data", "documents.json")

	var nodes []models.Node

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &nodes)
	if err != nil {
		return err
	}

	log.Info("Generating new elastic search client...")
	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		return err
	}
	log.Info("Connected to elastic search!")

	exists, err := client.IndexExists("ire").Do(context.Background())
	if err != nil {
		return err
	}

	if !exists {
		log.Warn("Index \"ire\" does not exist, please create it using curl. \n")
		log.Info("curl -XPUT 'http://localhost:9200/some_index' -H 'Content-Type: application/json' -d '{}' ")
		return nil
	}

	// Declare a new Bulk() object using the client instance
	bulk := client.Bulk()

	for _, node := range nodes {
		node.Timestamp = time.Now().Unix()
		req, err := prepareDocumentRequest(node)
		bulk.Add(&req)
		if err != nil {
			return err
		}
	}

	response, err := bulk.Do(context.Background())

	if err != nil {
		log.Fatal("bulk.Do(ctx) ERROR:", err)
	} else {
		indexed := response.Indexed()
		fmt.Println("nbulkResp.Indexed():", indexed)
		fmt.Println("bulkResp.Indexed() TYPE:", reflect.TypeOf(indexed))

		// Iterate over the bulkResp.Indexed() object returned from bulk.go
		t := reflect.TypeOf(indexed)
		fmt.Println("nt:", t)
		fmt.Println("NewBulkIndexRequest().NumberOfActions():", bulk.NumberOfActions())

		// Iterate over the document responses
		for i := 0; i < t.NumMethod(); i++ {
			method := t.Method(i)
			fmt.Println("nbulkResp.Indexed() METHOD NAME:", i, method.Name)
			fmt.Println("bulkResp.Indexed() method:", method)
		}

		// Return data on the documents indexed
		fmt.Println("nBulk response Index:", indexed)
		for _, info := range indexed {
			fmt.Println("nBulk response Index:", info)
			//fmt.Println("nBulk response Index:", info.Index)
		}
	}

	return nil
}

// Prepares a document index request which will
// be later added to a BulkRequest
func prepareDocumentRequest(node models.Node) (elastic.BulkIndexRequest, error) {
	req := elastic.NewBulkIndexRequest()
	req.OpType("index")
	req.Id(node.ID)
	req.Doc(node)
	req.Index("ire")
	return *req, nil
}
