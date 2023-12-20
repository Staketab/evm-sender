package vars

import (
	"log"
	"os"
	"time"
)

var (
	InfoLog        = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	ErrorLog       = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	Binary         = "evm-sender"
	Version        = "1.2.1"
	ConfigPath     = ".evm-sender/"
	ConfigFilePath = ".evm-sender/config.toml"
	SendBackTicker = time.NewTicker(2 * time.Second)
)
