package core

import (
	"flag"
	"github.com/andrepinto/holmes/config"
	"github.com/andrepinto/holmes/api"
)

type App struct {
	Api *api.Api
}

func InitApp()(*App){
	return &App{}
}

func Load() (*App, error){

	app, err := NewApp()

	if err != nil {
		return nil, err
	}

	var configFile string

	if !flag.Parsed() {
		flag.StringVar(&configFile, "config", "configuration.json",
			"JSON config or file:// path to JSON config file.")
		flag.Parse()
	}

	config, err := config.NewConfiguration(configFile)

	if err != nil {
		return nil, err
	}


	if(config.Api.Port>0){
		app.Api, _ = api.NewApi(config)
	}


	return app, nil
}

func(app *App) Run(){
	app.Api.Run()
}

func NewApp() (*App, error){
	app := InitApp()
	return app, nil
}

