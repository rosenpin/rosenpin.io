package models

import "net/http"

type handlerCreator interface {
	CreateHandler(configPath string) http.Handler
}
