package sonos

import (
	"errors"
	"log"
	"time"

	"github.com/szatmary/sonos"
)

// Sonos client
var Sonos = connect()

// Speakers available on local network
var Speakers = findSpeakers()

var (
	// ErrorRoomNotFound error
	ErrorRoomNotFound error = errors.New("room not found")
	// ErrorPlaylistNotFound error
	ErrorPlaylistNotFound error = errors.New("playlist not found")
)

func connect() *sonos.Sonos {
	s, err := sonos.NewSonos()
	if err != nil {
		log.Fatalln("Fatal: ", err.Error())
	}

	return s
}

// findSpeakers returns all zone speakers
func findSpeakers() []sonos.ZonePlayer {

	var players []sonos.ZonePlayer

	found, err := Sonos.Search()
	if err != nil {
		log.Fatalln("Fatal: ", err.Error())
	}
	to := time.After(10 * time.Second)

find:
	for {
		select {
		case <-to:
			break find
		case zp := <-found:
			players = append(players, *zp)
		}
	}

	return players
}

// GetSpeaker by name
func GetSpeaker(name string) (*sonos.ZonePlayer, error) {
	zp := &sonos.ZonePlayer{}
	for _, s := range Speakers {
		if s.RoomName() == name {
			zp = &s
		}
	}

	if zp == nil {
		return nil, ErrorRoomNotFound
	}

	return zp, nil
}
