package discovery

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)

type SherlockService struct {
	Name string `json:"name" yaml:"name"`
	Ip []string `json:"ip" yaml:",flow"`
}

type Environment struct {
	Services []SherlockService `json:"service" yaml:"services"`
}

const SHERLOCK_FILE  = "sherlock-services.yaml"

func LoadEnvironment() (*Environment, error){
	var data []byte
	var err error

	if data, err = ioutil.ReadFile(SHERLOCK_FILE); err != nil {
		return nil, fmt.Errorf("ERROR ON READ SHERLOCK SERVICES FILE: %s", err)
	}

	env := &Environment{}

	err = yaml.Unmarshal(data, &env)


	return env, nil
}