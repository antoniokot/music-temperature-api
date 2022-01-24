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
		ClientID:     "2c3a77eb3039418f86eaf972ebdd09e6",
		ClientSecret: "b37dd90e0d0741a08cd563a76928ad03",
		TokenURL:     spotify.TokenURL,
	}
	
	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("Erro ao tentar recuperar token de acesso: %v", err)
	}

	Client = spotify.Authenticator{}.NewClient(accessToken)
}