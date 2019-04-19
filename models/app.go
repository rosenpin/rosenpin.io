package models

// App represents a sub application to run under the server
type App struct {
	// ConfigPath is the path to the app specific configuration file
	ConfigPath string
	// Name is the name of the application, appropriatley
	Name string
	// Path is the path on the server where the app will be served on
	Path string
	// Creator is the creator function for the HTTP handler for the app
	Creator handlerCreator
}
