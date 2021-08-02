package config

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func SetupFirebase(account string) *auth.Client {

	serviceAccountFilePath, err := filepath.Abs(account)
	if err != nil {
		panic("Unable to load ServiceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountFilePath)

	// Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	// Firebase auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase auth error")
	}
	return auth
}
