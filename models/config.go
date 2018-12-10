package models

// Config is the structs that holds the app configurable data
type Config struct {
	ResourcesPath      string  `yaml:"ResourcesPath"`
	Port               float64 `yaml:"Port"`
	SSLCertificatePath string  `yaml:"SSLCertificateLocation"`
	BaseConfigPath     string  `yaml:"BaseConfigPath"`
}
