package diagnostic

import (
	"fmt"

	"github.com/andrepinto/holmes/discovery"
	"github.com/andrepinto/holmes/version"
)

type Diagnostic struct {
	Environment *discovery.Environment
}

type DiagnosticResult struct {
	Name string
}

type ServiceDiagnosticResult struct {
	Name string
	Status string
	Result bool
}

const (
	STATUS_OK = "OK"
	STATUS_ERROR = "ERROR"
)

func NewDiagnostic(env *discovery.Environment) (*Diagnostic, error) {

	dg := &Diagnostic{
		Environment: env,
	}

	return dg, nil
}

func (d *Diagnostic) Run() {
	fmt.Println("run diagnostic")

	for _, v := range d.Environment.Services {
		for _, ips := range v.Ip {
			result, _ := discovery.GetServiceInfo(ips)
			fmt.Println("---------------")
			fmt.Println(v.Name, ips)
			d.CheckDependency(result)

		}
	}
}

func (d *Diagnostic) GetServiceDetail(name string) *discovery.SherlockService {
	for _, v := range d.Environment.Services {
		if v.Name == name {
			return &v
		}
	}

	return nil
}

func (d *Diagnostic) CheckDependency(service *discovery.ServiceData) {
	for _, v := range service.Dependencies {
		sv := d.GetServiceDetail(v.Name)
		sdr := &ServiceDiagnosticResult{
			Name: sv.Name,
		}
		for _, ips := range sv.Ip {
			result, _ := discovery.GetServiceInfo(ips)
			isValid := version.CheckVersion(v.Version, result.Service.Version)
			sdr.Result = isValid
		}
		fmt.Println("%v",sdr )
	}


}