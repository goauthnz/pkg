package http

// Config is the configuration for the HTTP server.
type Config struct {
	// Port is the port to listen on.
	Port string `env:"PORT"`
}
