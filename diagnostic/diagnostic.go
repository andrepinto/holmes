package diagnostic

import (
	"fmt"

	"github.com/andrepinto/holmes/discovery"
)

type Diagnostic struct {
	Environment *discovery.Environment
}

type DiagnosticResult struct {
}

func NewDiagnostic(env *discovery.Environment) (*Diagnostic, error) {

	dg := &Diagnostic{
		Environment: env,
	}

	return dg, nil
}

func (d *Diagnostic) Run() {
	fmt.Println("run diagnostic")

	for _, v := range d.Environment.Services {
		fmt.Println(v.Name)
		for _, ips := range v.Ip {
			result, _ := discovery.GetServiceInfo(ips)
			fmt.Println(result.Service)
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
		fmt.Println(v)
		sv := d.GetServiceDetail(v.Name)
		for _, ips := range sv.Ip {
			result, _ := discovery.GetServiceInfo(ips)
			fmt.Println(result.Service.Version)
		}
	}
}