package main

import (
	"flag"
	"kimchi/api"
	"log"
)

func main() {
	// Parse Flag
	envPtr := flag.String("env", "dev", "deploy environment")
	flag.Parse()

	// Updater
	var updaterEnv string
	switch *envPtr {
	case "dev":
		updaterEnv = "./Redis.yaml"
	case "deploy":
		updaterEnv = "./Redis_deploy.yaml"
	}

	wsBase := api.New(updaterEnv)

	ws := wsBase.Serve("./Config.yaml", *envPtr)
	log.Fatal(ws.ListenAndServe())
}
