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
	Version        = "1.0.0"
	ConfigPath     = ".evm-sender/"
	ConfigFilePath = ".evm-sender/config.toml"
	MissingTicker  = time.NewTicker(20 * time.Second)
	//TpsTicker      = time.NewTicker(60 * time.Second)
)
