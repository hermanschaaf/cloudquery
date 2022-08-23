package client

import (
	v5 "github.com/heroku/heroku-go/v5"
	"net/http"
)

type Services struct {
	Heroku *v5.Service
}

func initServices(client *http.Client) (*Services, error) {
	h := v5.NewService(client)
	return &Services{
		Heroku: h,
	}, nil
}
