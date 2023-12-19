package main

import (
	"github.com/staketab/evm-sender/cmd/evm-sender/internal"
	"github.com/staketab/evm-sender/cmd/evm-sender/internal/var"
	"os"
)

func main() {
	if err := internal.RootCmd.Execute(); err != nil {
		vars.ErrorLog.Println(err)
		os.Exit(1)
	}
}
