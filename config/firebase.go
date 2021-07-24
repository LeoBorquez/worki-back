package config

import (
	"firebase.google.com/go/auth"
)

func setupFirebase() *auth.Client {
	return auth
}
