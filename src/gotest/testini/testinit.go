package main

import (
	"flag"
	"fmt"
	"github.com/larspensjo/config"
	"log"
)

var (
	configPath = "D:\\goproject\\goext\\src\\gotest\\testini\\config.ini"
	configFile = flag.String("configfile", configPath, "General configuration file")
)

//topic list
var TOPIC = make(map[string]string)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("topicArr") {
		if section, err := cfg.SectionOptions("topicArr"); err == nil {
			for _, v := range section {
				options, err := cfg.String("topicArr", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	//Initialized topic from the configuration END

	fmt.Println(TOPIC)
	fmt.Println(TOPIC["debug"])
}
