package config

// ServerConfig defines the configuration for the server
type ServerConfig struct {
	Port           int
}

// New return all constants using in Project such as Dialogflow's ProjectID, Line's ChannelID
func New() ServerConfig {
	// Set up the stripe key early on, so everything is nicely preconfigured
	// when it's used in other parts of the app.
	return ServerConfig{
		Port:           GetRequiredIntEnv("AUTH_PORT"),
	}
}
