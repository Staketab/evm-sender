package internal

import (
	"github.com/spf13/cobra"
	vars "github.com/staketab/evm-sender/cmd/internal/var"
	ff "github.com/staketab/evm-sender/cmd/pkg/func"
)

var RootCmd = &cobra.Command{
	Use:   vars.Binary,
	Short: "ETH Indexer",
	Long:  "ETH Indexer",
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior when the root command is executed
		vars.InfoLog.Println("Welcome to:", vars.Binary)
	},
}

func InitConfig() {
	err := ff.CreateDirectory(vars.ConfigPath)
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
}

func InitConfigFile() {
	err := ff.CreateConfigFile(vars.ConfigFilePath)
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
}

func ReadConfig() {
	config, err := ff.ReadConfigs()
	if err != nil {
		vars.ErrorLog.Println(err)
		return
	}
	if len(config.Default.Rpc) > 0 {
		vars.InfoLog.Println("Config read successfully!")
	} else {
		vars.InfoLog.Fatal("Config is empty.")
	}
}

func InitStart() {
	InitConfig()
	InitConfigFile()
	ReadConfig()
}
