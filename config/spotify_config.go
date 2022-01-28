package config

import (
	"context"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var SpotifyClient spotify.Client

func StartSpotifyConfig() {

	authConfig := &clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	
	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("Erro ao tentar recuperar token de acesso: %v", err)
	}

	SpotifyClient = spotify.Authenticator{}.NewClient(accessToken)
}