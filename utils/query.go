/*
 * Copyright (c) 2020.
 * Jim Filippou · jimfilippou8@gmail.com
 */

package utils

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/urfave/cli"
)

func Query(ctx *cli.Context, queriesPath string) (*elastic.SearchResult, error) {

	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		return nil, err
	}

	// TODO: Read queries file, iterate through queries & perform them
	matchQuery := elastic.NewMatchQuery("Text", "query")

	// This line is too JavaScripty, why the fuck did you make it like this?
	searchResult, err := client.Search().
		Index("ire").
		Query(matchQuery).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	return searchResult, nil
}
