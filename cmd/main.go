package main

import (
	"context"
	"fmt"
	"log"

	"github.com/39shin52/spotifyAPIParser/auth/app/auth"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to Load .env file: %v", err)
	}

	ctx := context.Background()
	client, err := auth.Auth(ctx)
	if err != nil {
		log.Fatalf("failed to auth: %v", err)
	}

	var albumId spotify.ID = "07w0rG5TETcyihsEIZR3qG"
	album, err := client.GetAlbum(ctx, albumId)
	if err != nil {
		log.Fatalf("failed to get album: %v", err)
	}

	for _, sa := range album.Tracks.Tracks {
		fmt.Println(sa.Name)
	}

	var trackID spotify.ID = "11dFghVXANMlKmJXsNCbNl"
	track, err := client.GetTrack(ctx, trackID)
	if err != nil {
		log.Fatalf("failed to get track: %v", err)
	}

	fmt.Println(track.Name)
}
