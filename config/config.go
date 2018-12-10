package config

import (
	"io/ioutil"

	"gitlab.com/rosenpin/rosenpin.io/models"
)

// Unmarshaler is the interface for objects that can
type Unmarshaler func(content []byte, target interface{}) error

// Loader loads the configuration from the configuration file into the configuration object
type Loader struct{ path string }

// NewLoader creates a new configuration loader object using the provided path
func NewLoader(path string) *Loader {
	return &Loader{path}
}

// Load loads the configuration from the configuration file into the configuration object
func (cl *Loader) Load(unmarshaler Unmarshaler, target *models.Config) error {
	content, err := ioutil.ReadFile(cl.path)
	if err != nil {
		return err
	}

	return unmarshaler(content, target)
}
