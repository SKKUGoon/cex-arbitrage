package main

import (
	"flag"
	"kimchi/api"
	"kimchi/dao"
	"log"
)

func main() {
	// Parse Flag
	envPtr := flag.String("env", "dev", "deploy environment")
	flag.Parse()

	// Updater
	var myEnv string
	switch *envPtr {
	case "dev":
		myEnv = "./config_dev.yaml"
	case "deploy":
		myEnv = "./config_deploy.yaml"
	}
	// PubSub
	dao.SignalMQ = dao.NewSignalReciever(myEnv)
	go func() {
		dao.SignalMQ.Run()
	}()

	wsBase := api.New(myEnv, *envPtr)

	ws := wsBase.Serve(myEnv, *envPtr)
	log.Fatal(ws.ListenAndServe())
}
