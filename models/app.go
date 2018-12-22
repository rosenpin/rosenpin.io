package models

// App represents a sub application to run under the server
type App struct {
	// ConfigPath is the path to the app specific configuration file
	ConfigPath string
	// Name is the name of the application, appropriatley, the path to the app on the server would be
	// for example 127.0.0.1/Name
	Name string
	// Creator is the creator function for the HTTP handler for the app
	Creator handlerCreator
}
