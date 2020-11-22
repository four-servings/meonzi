package config

import "os"

// Database database config
type Database struct{}

// User database config
func (d Database) User() string {
	if env := os.Getenv("DATABASE_USER"); env != "" {
		return env
	}
	return "root"
}

// Password database config
func (d Database) Password() string {
	if env := os.Getenv("DATABASE_PASSWORD"); env != "" {
		return env
	}
	return "test"
}

// Host database config
func (d Database) Host() string {
	if env := os.Getenv("DATABASE_HOST"); env != "" {
		return env
	}
	return "localhost"
}

// Port database config
func (d Database) Port() string {
	if env := os.Getenv("DATABASE_POST"); env != "" {
		return env
	}
	return "3306"
}

// Name database config
func (d Database) Name() string {
	if env := os.Getenv("DATABASE_NAME"); env != "" {
		return env
	}
	return "dropit"
}
