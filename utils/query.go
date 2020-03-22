/*
 * Copyright (c) 2020.
 * Jim Filippou · jimfilippou8@gmail.com
 */

package utils

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/urfave/cli"
	"io/ioutil"
	"strings"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// This function is devoted 100% for my friend Artemis
func fetchQueries(path string) ([]string, error) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var queries []string

	tokens := strings.Split(string(file), "\n")

	for i := 1; i <= len(tokens); i += 3 {
		queries = append(queries, tokens[i])
	}

	return queries, nil

}

func Query(ctx *cli.Context, queriesPath string) ([]*elastic.SearchResult, error) {

	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		return nil, err
	}

	queries, err := fetchQueries(queriesPath)
	if err != nil {
		return nil, err
	}

	var results []*elastic.SearchResult

	for _, query := range queries {

		matchQuery := elastic.NewMatchQuery("Text", query)

		// This line is too JavaScripty, why the fuck did you make it like this?
		searchResult, err := client.Search().
			Index("ire").
			Query(matchQuery).
			Pretty(true).
			Do(context.Background())
		if err != nil {
			return nil, err
		}

		results = append(results, searchResult)

	}

	return results, nil

}
