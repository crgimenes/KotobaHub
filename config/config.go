package config

import (
	"flag"
	"os"

	"golang.org/x/term"
)

type Config struct {
	ListemAddress string // Address to listen on
	DBPath        string // Path to the database file
	Debug         bool   // Enable debug mode
	IsTTY         bool   // Is a TTY
}

var (
	CFG = &Config{}
)

func Load() {
	CFG.DBPath = "file:data.db?mode=rwc&_journal_mode=WAL&_busy_timeout=10000`"

	if term.IsTerminal(int(os.Stdout.Fd())) {
		CFG.IsTTY = true
	}

	flag.StringVar(&CFG.DBPath, "db", CFG.DBPath, "Path to the database file")
	flag.BoolVar(&CFG.Debug, "debug", false, "Enable debug mode")
	flag.StringVar(&CFG.ListemAddress, "listen", ":8080", "Address to listen on")

	flag.Parse()
}
