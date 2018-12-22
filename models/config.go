package models

// Config is the structs that holds the app configurable data
type Config struct {
	ResourcesPath      string  `yaml:"resources_path"`
	StaticFilesPath    string  `yaml:"static_files"`
	Port               float64 `yaml:"port"`
	SSLCertificatePath string  `yaml:"ssl_certificate_location"`
	BaseConfigPath     string  `yaml:"configs_path"`
	UpgradeHTTP        bool    `yaml:"upgrade_http"`
}
