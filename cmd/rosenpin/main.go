package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"

	showcase "gitlab.com/rosenpin/git-project-showcaser/cmd/showcaser"
	"gitlab.com/rosenpin/rosenpin.io/config"
	"gitlab.com/rosenpin/rosenpin.io/models"
	yaml "gopkg.in/yaml.v2"
)

type handlerCreator interface {
	CreateHandler(configPath string) http.Handler
}

const (
	configFormat = "yml"
)

type app struct {
	configPath string
	name       string
	creator    handlerCreator
}

var (
	appsHandlers = []app{
		{configPath: "showcase.yml", name: "/", creator: showcase.NewProjectShowcase()},
	}
)

func main() {
	// Parse flags
	var configPath string

	flag.StringVar(&configPath, "c", "", "path to the configuration file")
	flag.Parse()

	config := loadConfig(configPath)

	startServers(config)
}

func loadConfig(configPath string) *models.Config {
	if configPath == "" {
		panic("no configuration file specified")
	}

	configLoader := config.NewLoader(configPath)

	config := &models.Config{}

	err := configLoader.Load(yaml.Unmarshal, config)
	if err != nil {
		panic(err)
	}

	validateConfig(config)

	return config
}

func validateConfig(config *models.Config) error {
	if config.Port > 65535 || config.Port <= 0 {
		return fmt.Errorf("invalid port number")
	}

	if _, err := os.Stat(config.ResourcesPath); err != nil {
		return fmt.Errorf("invalid resources path - %v", err)
	}

	return nil
}

func startServers(config *models.Config) {
	r := http.NewServeMux()
	for _, app := range appsHandlers {
		handler := app.creator.CreateHandler(path.Join(config.BaseConfigPath, app.configPath)).ServeHTTP
		r.HandleFunc(app.name, handler)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join(config.ResourcesPath, "static")))))
	http.Handle("/", r)

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", uint(config.Port)), nil); err != nil {
			panic(err)
		}
	}()

	err := http.ListenAndServeTLS(":443", path.Join(config.SSLCertificatePath, "fullchain.pem"), path.Join(config.SSLCertificatePath, "privkey.pem"), nil)
	if err != nil {
		fmt.Println(err)
	}
	blocker := make(chan int)
	<-blocker
}
