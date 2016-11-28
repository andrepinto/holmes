package config

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/andrepinto/holmes/util"
)

type Configuration struct {
	Service Service `json:"service,omitempty"`
	Api Api  `json:"api,omitempty"`
}



func NewConfiguration(path string)(*Configuration, error){
	var data []byte
	var err error

	if data, err = ioutil.ReadFile(path); err != nil {
		return nil, fmt.Errorf("ERROR ON READ CONFIG FILE: %s", err)
	}


	template, err := ApplyTemplate(data)
	if err != nil {
		return nil, fmt.Errorf(
			"ERROR ON READ TEMPLATE: %v", err)
	}

	configuration, err := unmarshalConfig(template)

	if err != nil {
		return nil, fmt.Errorf(
			" %v", err)
	}


	return configuration, nil
}

func unmarshalConfig(data []byte) (*Configuration, error) {
	conf := Configuration{}
	if err := json.Unmarshal(data, &conf); err != nil {
		syntax, ok := err.(*json.SyntaxError)
		if !ok {
			return &conf, fmt.Errorf(
				"Could not parse configuration: %s",
				err)
		}
		return nil, util.NewJSONparseError(data, syntax)
	}

	return &conf, nil
}

