package main

import (
	"context"
	"fmt"
	"log"

	"github.com/39shin52/spotifyAPIParser/auth/app/auth"
	"github.com/39shin52/spotifyAPIParser/auth/app/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to Load .env file: %v", err)
	}

	var input string
	fmt.Printf("Input Trakc name: ")
	fmt.Scan(&input)

	ctx := context.Background()
	client, err := auth.Auth(ctx)
	if err != nil {
		log.Fatalf("failed to auth: %v", err)
	}
	c := handler.Client{Client: client}

	if err := c.SearchTrack(ctx, input); err != nil {
		log.Fatalf(err.Error())
	} else {
		fmt.Println("ok")
	}

	tempo := c.GetTrackBpm(ctx, input)
	fmt.Println(tempo)
}
