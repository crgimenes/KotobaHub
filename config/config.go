package config

import "flag"

type Config struct {
	ListemAddress string // Address to listen on
	DBPath        string // Path to the database file
	Debug         bool   // Enable debug mode
}

var (
	CFG = &Config{}
)

func Load() {
	CFG.DBPath = "file:data.db?mode=rwc&_journal_mode=WAL&_busy_timeout=10000`"

	flag.StringVar(&CFG.DBPath, "db", CFG.DBPath, "Path to the database file")
	flag.BoolVar(&CFG.Debug, "debug", false, "Enable debug mode")
	flag.StringVar(&CFG.ListemAddress, "listen", ":8080", "Address to listen on")

	flag.Parse()
}
