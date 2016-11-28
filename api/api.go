package api

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/andrepinto/holmes/config"
	"encoding/json"
	"fmt"
	"github.com/andrepinto/holmes/discovery"
	"github.com/andrepinto/holmes/diagnostic"
)

type Api struct {
	Port int
	Configuration *config.Configuration
	router *mux.Router
}

type StatusResponse struct {
	Status string `json:"status"`
}



func NewApi(config *config.Configuration) (*Api, error){
	return &Api{
		Port:config.Api.Port,
		Configuration: config,
		router: mux.NewRouter(),
	}, nil
}



func(api *Api) statusHandler(w http.ResponseWriter, r *http.Request) {
	status := StatusResponse{"ok"}
	json.NewEncoder(w).Encode(status)
}

func(api *Api) envHandler(w http.ResponseWriter, r *http.Request) {
	env, _ := discovery.LoadEnvironment()
	json.NewEncoder(w).Encode(env)
}

func(api *Api) diagnosticHandler(w http.ResponseWriter, r *http.Request) {
	env, _ := discovery.LoadEnvironment()
	dg, _ := diagnostic.NewDiagnostic(env)
	dg.Run()
	json.NewEncoder(w).Encode(dg)
}

func(api *Api) Run() {

	api.router.HandleFunc("/status", api.statusHandler)
	api.router.HandleFunc("/env", api.envHandler)
	api.router.HandleFunc("/diagnostic", api.diagnosticHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",api.Port), api.router))
}
