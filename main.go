package main

import (
	"log"

	"github.com/mikejwhitehead/jukebox/config"
	"github.com/mikejwhitehead/jukebox/sonos"
	"github.com/tarm/serial"
)

func main() {

	// Load config from file
	cfg, err := config.Load("./config.yaml")
	if err != nil {
		log.Fatalln("Fatal: ", err.Error())
	}

	_, err = sonos.GetSpeaker(cfg.Room)
	if err != nil {
		log.Fatalln("Fatal: ", err.Error())
	}

	c := &serial.Config{Name: cfg.InputDevice, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
			log.Fatal(err)
	}

	buf := make([]byte, 128)
	_, err = s.Read(buf)
	if err != nil {
			log.Fatal(err)
	}

	log.Println(string(buf))

}