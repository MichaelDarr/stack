package config

import "fmt"

// PostgresConfig contains PG database credentials.
type PostgresConfig struct {
	Debug             bool
	Host              string
	Name              string
	Password          string
	Port              int
	Schema            string
	User              string
	MigrationFilepath string
	SSLEnabled        bool
}

// DSN constructs a postgres DSN.
func (cfg PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=America/New_York",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
		sslMode(cfg.SSLEnabled),
	)
}

// URL constructs a postgres connection URL.
func (cfg PostgresConfig) URL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		sslMode(cfg.SSLEnabled),
	)
}

// sslMode returns an enabled or disabled ssl connection mode string.
func sslMode(enabled bool) string {
	if enabled {
		return "enable"
	}
	return "disable"
}
