package main

import (
	"fmt"
	"github.com/zmb3/spotify/v2"
	"log"
	"strings"
)

func CreatePlaylist(client *spotify.Client, playlistName string, userId string, tracks []spotify.SimpleTrack) {
	fmt.Println("Creating playlist:", playlistName)
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

func PlaylistContainTrack(playlist []spotify.PlaylistItemTrack, track spotify.PlaylistItemTrack) bool {
	for _, v := range playlist {
		if v.Track.Name == track.Track.Name {
			return true
		}
	}

	return false
}

func IntersectPlaylists(a []spotify.PlaylistItemTrack, b []spotify.PlaylistItemTrack) []spotify.SimpleTrack {
	set := make([]spotify.SimpleTrack, 0)

	for _, v := range a {
		if PlaylistContainTrack(b, v) {
			set = append(set, v.Track.SimpleTrack)
		}
	}

	return set
}

func GetUserPlaylistTracksWithPrefix(client *spotify.Client, prefix string) map[string][]spotify.PlaylistItemTrack {
	userPlaylists, err := client.CurrentUsersPlaylists(ctx)

	if err != nil {
		log.Fatal("error getting user playlists", err.Error())
		return nil
	}

	playlists := make(map[string][]spotify.PlaylistItemTrack)

	for playlistPage := 1; ; playlistPage++ {
		for _, playlist := range userPlaylists.Playlists {
			if strings.HasPrefix(playlist.Name, prefix) {
				playlistTracks, _ := client.GetPlaylistItems(ctx, playlist.ID)
				var tracks []spotify.PlaylistItemTrack
				for trackPage := 1; ; trackPage++ {
					for _, track := range playlistTracks.Items {
						tracks = append(tracks, track.Track)
					}

					err = client.NextPage(ctx, playlistTracks)
					if err == spotify.ErrNoMorePages {
						break
					}
					if err != nil {
						log.Fatal(err)
					}
				}

				playlists[playlist.Name] = tracks
			}

		}

		err = client.NextPage(ctx, userPlaylists)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	return playlists
}

func CreateIntersectingPlaylists(client *spotify.Client, userId string) {
	playlists := GetUserPlaylistTracksWithPrefix(client, "Library")

	positiveInstrumental := IntersectPlaylists(playlists["Library - Positive"], playlists["Library - Instrumental"])
	CreatePlaylist(client, "Library - Positive Instrumentals", userId, positiveInstrumental)

	negativeInstrumental := IntersectPlaylists(playlists["Library - Negative"], playlists["Library - Instrumental"])
	CreatePlaylist(client, "Library - Negative Instrumentals", userId, negativeInstrumental)

	positiveAcoustic := IntersectPlaylists(playlists["Library - Positive"], playlists["Library - Acoustic"])
	CreatePlaylist(client, "Library - Positive Acoustic", userId, positiveAcoustic)

	negativeAcoustic := IntersectPlaylists(playlists["Library - Negative"], playlists["Library - Acoustic"])
	CreatePlaylist(client, "Library - Negative Acoustic", userId, negativeAcoustic)

	highEnergyInstrumental := IntersectPlaylists(playlists["Library - High Energy"], playlists["Library - Instrumental"])
	CreatePlaylist(client, "Library - HighEnergyInstrumental", userId, highEnergyInstrumental)

	lowEnergyInstrumental := IntersectPlaylists(playlists["Library - Low Energy"], playlists["Library - Instrumental"])
	CreatePlaylist(client, "Library - Low Energy Instrumental", userId, lowEnergyInstrumental)

}
