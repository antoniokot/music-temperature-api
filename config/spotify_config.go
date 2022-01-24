package config

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var Client spotify.Client

func StartSpotifyConfig() {

	authConfig := &clientcredentials.Config{
		ClientID:     "08c1a6be652e4fdca07f1815bfd167e4",
		ClientSecret: "<your_client_secret",
		TokenURL:     spotify.TokenURL,
	}
	
	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("Erro ao tentar recuperar token de acesso: %v", err)
	}

	Client = spotify.Authenticator{}.NewClient(accessToken)
}