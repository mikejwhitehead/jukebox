package main

import (
	"fmt"

	"github.com/mikejwhitehead/jukebox/config"
	"github.com/mikejwhitehead/jukebox/ent/card"
	"github.com/mikejwhitehead/jukebox/sonos"
	"github.com/tarm/serial"

	"context"
	"log"

	"github.com/mikejwhitehead/jukebox/ent"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
}

func listen(input string, output chan string) {
	c := &serial.Config{Name: input, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 256)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		cardID := string(buf[:n])

		output <- cardID
	}

}

func process(ctx context.Context, client *ent.Client, cardID string) {
	log.Println(cardID)

	c, err := client.Card.
		Query().
		Where(card.NameEQ(cardID)).
		Only(ctx)
	if err != nil {
		log.Println(cardID, " not found")
		return
	}

	log.Println(cardID, " playlist set to ", c.Edges.Playlist.Name)

}

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

	output := make(chan string)
	go listen(cfg.InputDevice, output)

	for id := range output {
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
