/*
 * Copyright (c) 2020.
 * Jim Filippou · jimfilippou8@gmail.com
 */

package models

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	ElasticSearchInstance string
	KibanaInstance        string
	path                  string
}

func NewConfiguration(path string) (*Configuration, error) {
	var conf Configuration
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
