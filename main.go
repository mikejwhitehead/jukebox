package main

import (
	"fmt"
	"log"

	"github.com/mikejwhitehead/jukebox/config"
	"github.com/mikejwhitehead/jukebox/sonos"
	musicservices "github.com/szatmary/sonos/MusicServices"
)

func main() {

	// Load config from file
	cfg, err := config.Load("./config.yaml")
	if err != nil {
		log.Fatalln("Fatal: ", err.Error())
	}

	zp, err := sonos.GetSpeaker(cfg.Room)
	if err != nil {
		log.Fatalln("Fatal: ", err.Error())
	}

	listArgs := musicservices.ListAvailableServicesArgs{}
	svc, err := zp.MusicServices.ListAvailableServices(zp.HttpClient, &listArgs)
	if err != nil {
		log.Fatalln("Fatal: ", err.Error())
	}

	fmt.Println(svc.AvailableServiceTypeList)

	// if err := zp.Play(); err != nil {
	// 	log.Fatalln("Fatal: ", err.Error())
	// }

}