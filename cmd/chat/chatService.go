package main

import (
	"all/internal/config"
	"all/internal/config/service/chat"
	"flag"
	"fmt"
	"os"
)

func main() {

	cfg := readConfig()
	//le injecto a el chat service una configuracion
	service, _ := chat.New(cfg)
	fmt.Println(cfg.Version)
	fmt.Println(cfg.DB.Driver)

	for _, m := range service.FindAll() {
		fmt.Println(m)
	}
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config/config.yaml", "sarasa")
	flag.Parse()
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
