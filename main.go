package main

import (
	"runtime"
	"github.com/andrepinto/holmes/core"
	"log"

)

func main(){
	runtime.GOMAXPROCS(1)

	app, configErr := core.Load()
	if configErr != nil {
		log.Fatal(configErr)
	}


	app.Run()
}



