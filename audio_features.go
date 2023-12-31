package main

import (
	"github.com/zmb3/spotify/v2"
	"log"
)

func GetTrackFeatures(client *spotify.Client, track spotify.FullTrack) *spotify.AudioFeatures {
	features, err := client.GetAudioFeatures(ctx, track.ID)
	if err != nil {
		log.Fatal("error getting audio features...", err.Error())
		return nil
	}
	return features[0]
}

func IsTrackAcoustic(trackFeatures *spotify.AudioFeatures) bool {
	return trackFeatures.Acousticness >= 0.7
}

func IsTrackHighEnergy(trackFeatures *spotify.AudioFeatures) bool {
	return trackFeatures.Energy >= 0.7
}

func IsTrackLowEnergy(trackFeatures *spotify.AudioFeatures) bool {
	return trackFeatures.Energy <= 0.4
}

func IsTrackInstrumental(trackFeatures *spotify.AudioFeatures) bool {
	return trackFeatures.Instrumentalness >= 0.7
}

func IsTrackPositive(trackFeatures *spotify.AudioFeatures) bool {
	return trackFeatures.Valence >= 0.7
}

func IsTrackNegative(trackFeatures *spotify.AudioFeatures) bool {
	return trackFeatures.Valence <= 0.4
}

func CreateFeaturePlaylists(client *spotify.Client, user *spotify.PrivateUser) {
	currentUserTracks, err := client.CurrentUsersTracks(ctx)
	if err != nil {
		log.Fatal("error getting user tracks...", err.Error())
		return
	}

	log.Printf("Playlist has %d total tracks", currentUserTracks.Total)
	var acousticTracks []spotify.SimpleTrack
	var highEnergyTracks []spotify.SimpleTrack
	var lowEnergyTracks []spotify.SimpleTrack
	var instrumentalTracks []spotify.SimpleTrack
	var positiveTracks []spotify.SimpleTrack
	var negativeTracks []spotify.SimpleTrack

	for page := 1; ; page++ {
		for _, track := range currentUserTracks.Tracks {
			//fmt.Println(track.Artists[0].Name, track.Name)
			trackFeatures := GetTrackFeatures(client, track.FullTrack)
			if IsTrackAcoustic(trackFeatures) {
				acousticTracks = append(acousticTracks, track.SimpleTrack)
			}
			if IsTrackHighEnergy(trackFeatures) {
				highEnergyTracks = append(highEnergyTracks, track.SimpleTrack)
			}
			if IsTrackLowEnergy(trackFeatures) {
				lowEnergyTracks = append(lowEnergyTracks, track.SimpleTrack)
			}
			if IsTrackInstrumental(trackFeatures) {
				instrumentalTracks = append(instrumentalTracks, track.SimpleTrack)
			}
			if IsTrackPositive(trackFeatures) {
				positiveTracks = append(positiveTracks, track.SimpleTrack)
			}
			if IsTrackNegative(trackFeatures) {
				negativeTracks = append(negativeTracks, track.SimpleTrack)
			}
		}

		err = client.NextPage(ctx, currentUserTracks)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	CreatePlaylist(client, "Library - Acoustic", user.ID, acousticTracks)
	CreatePlaylist(client, "Library - High Energy", user.ID, highEnergyTracks)
	CreatePlaylist(client, "Library - Low Energy", user.ID, lowEnergyTracks)
	CreatePlaylist(client, "Library - Instrumental", user.ID, instrumentalTracks)
	CreatePlaylist(client, "Library - Positive", user.ID, positiveTracks)
	CreatePlaylist(client, "Library - Negative", user.ID, negativeTracks)
}
