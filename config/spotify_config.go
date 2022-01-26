package config

import (
	"context"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"github.com/joho/godotenv"
)

var Client spotify.Client

func StartSpotifyConfig() {

	err := godotenv.Load(".env")

	if err != nil {
    log.Fatalf("Erro ao carregar arquivo .env")
  }

	authConfig := &clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	
	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("Erro ao tentar recuperar token de acesso: %v", err)
	}

	Client = spotify.Authenticator{}.NewClient(accessToken)
}