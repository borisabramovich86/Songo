package main

import (
	"context"
	"fmt"
	"songo/auth"
	"songo/playlists"
)

// func Search(ctx context.Context, client *spotify.Client, searchTerm string) []spotify.SimpleTrack {
// 	results, err := client.Search(ctx, searchTerm, spotify.SearchTypeTrack)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var list2 []spotify.SimpleTrack
// 	for _, x := range results.Tracks.Tracks {
// 		list2 = append(list2, x.SimpleTrack) // note the = instead of :=
// 	}
// 	return list2
// }

// func CreateRandomPlaylist(client *spotify.Client, userId string) {
// 	babbler := babble.NewBabbler()
// 	babbler.Count = 1

// 	for count := 1; count <= 10; count++ {
// 		randomWord := babbler.Babble()
// 		tracks := Search(client, randomWord)
// 		CreatePlaylist(client, "randPlaylist - "+randomWord, userId, tracks)
// 	}
// }

func main() {
	ctx := context.Background()
	client, user := auth.Authenticate()
	playlists.CreateFeaturePlaylists(ctx, client, user)
	// CreateIntersectingPlaylists(client, user.ID)
	// CreateRandomPlaylist(client, user.ID)

	fmt.Println("Done")
}
