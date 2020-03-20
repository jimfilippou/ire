/*
 * Copyright (c) 2020.
 * Jim Filippou · jimfilippou8@gmail.com
 */

package utils

import "github.com/olivere/elastic/v7"

func Query(q string) (interface{}, error) {

	_, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
