package config

// ServerConfig defines the configuration for the server
type ServerConfig struct {
	GQLPath        string
	PlaygroundPath string
	Port           int
	Postgres       PostgresConfig
}

// New return all constants using in Project such as Dialogflow's ProjectID, Line's ChannelID
func New() ServerConfig {
	// Set up the stripe key early on, so everything is nicely preconfigured
	// when it's used in other parts of the app.
	return ServerConfig{
		GQLPath:        GetRequiredEnv("BACKEND_GQL_PATH"),
		PlaygroundPath: GetRequiredEnv("BACKEND_PLAYGROUND_PATH"),
		Port:           GetRequiredIntEnv("BACKEND_PORT"),
		Postgres: PostgresConfig{
			Debug:             GetRequiredBoolEnv("BACKEND_POSTGRES_DEBUG"),
			Host:              GetRequiredEnv("POSTGRES_HOST"),
			Name:              GetRequiredEnv("POSTGRES_DB"),
			Password:          GetRequiredEnv("POSTGRES_PASSWORD"),
			Port:              GetRequiredIntEnv("POSTGRES_PORT"),
			Schema:            GetRequiredEnv("POSTGRES_SCHEMA"),
			User:              GetRequiredEnv("POSTGRES_USER"),
			MigrationFilepath: GetRequiredEnv("MIGRATION_FILEPATH"),
			SSLEnabled:        GetBoolEnv("POSTGRES_SSL_ENABLED", false),
		},
	}
}
