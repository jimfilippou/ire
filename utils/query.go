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

func Query(ctx *cli.Context, query string) (*elastic.SearchResult, error) {

	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		return nil, err
	}

	matchQuery := elastic.NewMatchQuery("Text", query)

	searchResult, err := client.Search().
		Index("ire").            // search in index "twitter"
		Query(matchQuery).       // specify the query
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute

	if err != nil {
		return nil, err
	}

	return searchResult, nil
}
