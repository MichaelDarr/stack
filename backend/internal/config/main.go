package config

// PostgresConfig contains PG database credentials
type PostgresConfig struct {
	Debug    bool
	Host     string
	Name     string
	Password string
	Port     int
	Schema   string
	User     string
}

// ServerConfig defines the configuration for the server
type ServerConfig struct {
	Host           string
	GQLPath        string
	PlaygroundPath string
	Port           int
	Scheme         string
	Postgres       PostgresConfig
}

// New return all constants using in Project such as Dialogflow's ProjectID, Line's ChannelID
func New() ServerConfig {
	// Set up the stripe key early on, so everything is nicely preconfigured
	// when it's used in other parts of the app.
	return ServerConfig{
		Host:           GetRequiredEnv("BACKEND_HOST"),
		GQLPath:        GetRequiredEnv("BACKEND_GQL_PATH"),
		PlaygroundPath: GetRequiredEnv("BACKEND_PLAYGROUND_PATH"),
		Port:           GetRequiredIntEnv("BACKEND_PORT"),
		Scheme:         GetRequiredEnv("BACKEND_SCHEME"),
		Postgres: PostgresConfig{
			Debug:    GetRequiredBoolEnv("BACKEND_POSTGRES_DEBUG"),
			Host:     GetRequiredEnv("POSTGRES_HOST"),
			Name:     GetRequiredEnv("POSTGRES_DB"),
			Password: GetRequiredEnv("POSTGRES_PASSWORD"),
			Port:     GetRequiredIntEnv("POSTGRES_PORT"),
			Schema:   GetRequiredEnv("POSTGRES_SCHEMA"),
			User:     GetRequiredEnv("POSTGRES_USER"),
		},
	}
}
