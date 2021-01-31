package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/mikejwhitehead/jukebox/config"
	"github.com/mikejwhitehead/jukebox/ent/card"
	"github.com/mikejwhitehead/jukebox/sonos"
	"github.com/tarm/serial"

	"context"

	"github.com/mikejwhitehead/jukebox/ent"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func listen(device string) chan string {
	c := &serial.Config{Name: device, Baud: 9600}
	log.Infoln("connecting to serial port ", device)
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	input := make(chan string)
	go func(chan string) {
		for {
			buf := make([]byte, 256)
			n, err := s.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			scan := string(buf[:n])
			cardID := strings.Replace(strings.TrimSpace(scan), "Tap card key : ", "", -1)
			log.WithFields(log.Fields{
				"Card":   cardID,
				"Device": device,
			}).Infoln("card read from scan")
			input <- cardID
		}
	}(input)

	log.WithFields(log.Fields{
		"Device": device,
	}).Infoln("listener started")

	return input
}

func process(ctx context.Context, client *ent.Client, cardID string) {
	entry := log.WithFields(log.Fields{
		"Card": cardID,
	})

	entry.Infof("finding playlist for card")
	c, err := client.Card.
		Query().
		Where(card.NameEQ(cardID)).
		Only(ctx)
	if err != nil {
		entry.Warnln("card playlist not found")
		return
	}

	entry.Data["Playlist"] = c.Edges.Playlist.Name
	entry.Infoln("card playlist found")
}

func main() {
	printFiglet()
	// Load config from file
	log.Infoln("reading from ./config.yaml")
	cfg, err := config.Load("./config.yaml")
	if err != nil {
		log.Fatalln("Fatal: ", err.Error())
	}
	log.Info("config loaded successfully")

	_, err = sonos.GetSpeaker(cfg.Room)
	if err != nil {
		log.Fatalln("Fatal: ", err.Error())
	}

	client, err := ent.Open("sqlite3", "file:ent.db?mode=rwc&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := listen(cfg.InputDevice)

	for id := range in {
		process(ctx, client, id)
	}

	// plist, err := createPlaylist(ctx, client, "tonights music 2", "http://play.list")
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// cardName := "4D:A2:50:23"

	// card, err := createCard(ctx, client, cardName, plist.ID)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }

	// _, err = createCardScan(ctx, client, card.ID, plist.ID)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }

}

func createPlaylist(ctx context.Context, client *ent.Client, name, url string) (*ent.Playlist, error) {
	p, err := client.Playlist.
		Create().
		SetName(name).
		SetMediaURL(url).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create playlist: %v", err)
	}
	log.Println("playlist was created: ", p)

	return p, nil
}

func createCard(ctx context.Context, client *ent.Client, name string, plist int) (*ent.Card, error) {
	c, err := client.Card.
		Create().
		SetName(name).
		SetPlaylistID(plist).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create card: %v", err)
	}

	log.Println("card was created: ", c)
	return c, nil
}

func createCardScan(ctx context.Context, client *ent.Client, card, plist int) (*ent.CardScan, error) {
	c, err := client.CardScan.
		Create().
		SetCardID(card).
		SetPlaylistID(plist).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create card scan: %v", err)
	}

	log.Println("card scan was created: ", c)
	return c, nil
}

func printFiglet() {

	colorPurple := "\033[35m"

	figlet := `
*********************************************
   _       _        _
  (_)_   _| | _____| |__   _____  __
  | | | | | |/ / _ | '_ \ / _ \ \/ /
  | | |_| |   |  __| |_) | (_) >  <
 _/ |\__,_|_|\_\___|_.__/ \___/_/\_\
|__/
	  
*********************************************`

	fmt.Println(string(colorPurple), figlet)

	return
}
