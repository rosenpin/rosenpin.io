package handler

import (
	"net/http"

	config_manager "gitlab.com/rosenpin/config-manager"
	yaml "gopkg.in/yaml.v2"
)

type rosenpinHandlerCreator struct {
}

func NewRosenpinHandlerCreator() *rosenpinHandlerCreator {
	return &rosenpinHandlerCreator{}
}

func (r *rosenpinHandlerCreator) CreateHandler(configPath string) http.Handler {
	configLoader := config_manager.NewLoader(configPath)
	config := config{}

	err := configLoader.Load(yaml.Unmarshal, &config)
	if err != nil {
		panic(err)
	}

	return newRosenpinHandler(config.FilesPath)
}

type handler struct {
	handler http.Handler
}

func newRosenpinHandler(fileServerPath string) *handler {
	return &handler{handler: http.FileServer(http.Dir(fileServerPath))}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
