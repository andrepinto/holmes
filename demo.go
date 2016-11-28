package main

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"github.com/andrepinto/holmes/discovery"
)

type T struct {
	A string
	B struct {
		  RenamedC int   `yaml:"c"`
		  D        []int `yaml:",flow"`
	  }
}

func main()  {

	var data []byte
	var err error

	if data, err = ioutil.ReadFile("serlock.yaml"); err != nil {
		fmt.Println(err)
	}

	env := &discovery.Environment{}

	err = yaml.Unmarshal(data, &env)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", env)

	s1:=discovery.SherlockService{
		Name:"sales",
		Ip:[]string{"1","2"},
	}
	s2:=discovery.SherlockService{
		Name:"sales2",
		Ip:[]string{"1","2"},
	}

	env2 := &discovery.Environment{
		Services:[]discovery.SherlockService{s1,s2},
	}


	d, err := yaml.Marshal(&env2)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
}
