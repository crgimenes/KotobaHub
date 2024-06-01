package config

import "flag"

type Config struct {
	DBPath string // Path to the database file
	Debug  bool   // Enable debug mode
}

var (
	CFG = &Config{}
)

func New() *Config {
	CFG.DBPath = "data.db"

	flag.StringVar(&CFG.DBPath, "db", CFG.DBPath, "Path to the database file")
	flag.BoolVar(&CFG.Debug, "debug", false, "Enable debug mode")

	flag.Parse()

	return CFG
}
