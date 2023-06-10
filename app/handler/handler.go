package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify/v2"
)

type Client struct {
	*spotify.Client
}

type Handler interface {
	SearchTrack(ctx context.Context, trackName string) error
	GetTrack(ctx context.Context, trackName string) spotify.SimpleTrack
	GetTrackBpm(ctx context.Context, trackName string) int
}

func (c *Client) SearchTrack(ctx context.Context, trackName string) error {
	if len(trackName) == 0 {
		return fmt.Errorf("track name is empty")
	}
	_, err := c.Search(ctx, trackName, spotify.SearchTypeTrack)
	if err != nil {
		return fmt.Errorf("failed to search the track: %v", err)
	}

	return nil
}

func (c *Client) GetTrack(ctx context.Context, trackName string) spotify.SimpleTrack {
	panic("implement!")
}

func (c *Client) GetTrackBpm(ctx context.Context, trackName string) int {
	if len(trackName) == 0 {
		log.Fatalf("track name is empty")
		return 0
	}

	track, err := c.Search(ctx, trackName, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatalf("failed to search track: %v", err)
		return 0
	}

	fmt.Println("found track name:", track.Tracks.Tracks[0].Name)
	trackID := track.Tracks.Tracks[0].SimpleTrack.ID // first hit track's id

	features, err := c.GetAudioFeatures(ctx, trackID)
	if err != nil {
		log.Fatalf("failed to get track features: %v", err)
	}

	var tempo int
	for _, af := range features {
		tempo = int(af.Tempo)
	}

	return tempo
}
