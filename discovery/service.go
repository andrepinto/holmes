package discovery

import (
	"github.com/parnurzeal/gorequest"
	"fmt"
	"encoding/json"
	"github.com/andrepinto/holmes/version"
)


type ServiceData struct {
	Service *Service `json:"service,omitempty"`
	Dependencies *[]Dependency  `json:"dependencies,omitempty"`
}

func(service *ServiceData) VersionValidator(v string)  {
	version.CheckVersion("","")
}

func GetServiceInfo(endpoint string)(*ServiceData, []error){
	request := gorequest.New()

	uri := fmt.Sprintf("http://%s/info", endpoint)

	_, body, err := request.Get(uri).End()

	if err != nil {
		return nil, err
	}

	msg := []byte(body)

	serviceData := &ServiceData{}
	json.Unmarshal(msg, serviceData)

	return serviceData, nil
}


