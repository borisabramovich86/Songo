package main

import (
	"fmt"
	"github.com/zmb3/spotify/v2"
	"log"
)

func Search(client *spotify.Client, searchTerm string) {
	results, err := client.Search(ctx, searchTerm, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	track := results.Tracks.Tracks[0]
	fmt.Printf("Track is %s by %s", track.Name, track.Artists[0].Name)
}

func main() {
	client, user := Authenticate()
	CreateFeaturePlaylists(client, user)
	CreateIntersectingPlaylists(client, user.ID)
	fmt.Println("Done")
}
