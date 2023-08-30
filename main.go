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

func CreatePlaylist(client *spotify.Client, playlistName string, userId string, tracks []spotify.FullTrack) {
	playlist, err := client.CreatePlaylistForUser(ctx, userId, playlistName, "", false, false)
	if err != nil {
		log.Fatal("error creating playlist", err.Error())
		return
	}
	for _, acousticTrack := range tracks {
		_, err := client.AddTracksToPlaylist(ctx, playlist.ID, acousticTrack.ID)
		if err != nil {
			log.Fatal("error Adding track to playlist", err.Error())
			return
		}
	}
}

func main() {
	client, user := Authenticate()
	CreateFeaturePlaylists(client, user)
}
