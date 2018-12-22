package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	config "gitlab.com/rosenpin/config-manager"
	showcase "gitlab.com/rosenpin/git-project-showcaser/cmd/showcaser"
	"gitlab.com/rosenpin/rosenpin.io/models"
	yaml "gopkg.in/yaml.v2"
)

var (
	appsHandlers = []models.App{
		{ConfigPath: "showcase.yml", Name: "/", Creator: showcase.NewProjectShowcase()},
	}
)

func main() {
	var configPath string

	// Parse flags
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

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}

	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		http.StatusTemporaryRedirect)
}

func startServers(config *models.Config) {
	mux := http.NewServeMux()
	// Load HTTP handlers
	for _, app := range appsHandlers {
		handler := app.Creator.CreateHandler(path.Join(config.BaseConfigPath, app.ConfigPath)).ServeHTTP
		mux.HandleFunc(app.Name, handler)
	}

	// Handle static files
	http.Handle(config.StaticFilesPath, http.StripPrefix(config.StaticFilesPath, http.FileServer(http.Dir(path.Join(config.ResourcesPath, config.StaticFilesPath)))))
	// Handle created applications
	http.Handle("/", mux)

	go func() {
		err := http.ListenAndServeTLS(":443", path.Join(config.SSLCertificatePath, "fullchain.pem"), path.Join(config.SSLCertificatePath, "privkey.pem"), nil)
		if err != nil {
			fmt.Println(err)
		}
	}()

	if config.UpgradeHTTP {
		if err := http.ListenAndServe(fmt.Sprintf(":%.0f", config.Port), http.HandlerFunc(redirect)); err != nil {
			panic(err)
		}
	} else {

		if err := http.ListenAndServe(fmt.Sprintf(":%.0f", config.Port), nil); err != nil {
			panic(err)
		}
	}

	// Block exit
	blocker := make(chan int)
	<-blocker
}
