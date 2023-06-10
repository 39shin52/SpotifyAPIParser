package handler

import (
	"context"
	"fmt"

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
	panic("implement!")
}
