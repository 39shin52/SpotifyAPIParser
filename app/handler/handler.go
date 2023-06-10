package handler

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

type Client struct {
	client *spotify.Client
}

type Handler interface {
	SearchTrack(ctx context.Context, trackName string) error
	GetTrack(ctx context.Context, trackName string) spotify.SimpleTrack
	GetTrackBpm(ctx context.Context, trackName string) int
}

func (c Client) SearchTrack(ctx context.Context, trackName string) error {
	panic("implement!")
}

func (c Client) GetTrack(ctx context.Context, trackName string) spotify.SimpleTrack {
	panic("implement!")
}

func (c Client) GetTrackBpm(ctx context.Context, trackName string) int {
	panic("implement!")
}
