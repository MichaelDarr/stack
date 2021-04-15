package config

import "fmt"

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

// DSN constructs a postgres DSN
func (cfg *PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/New_York",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
	)
}
