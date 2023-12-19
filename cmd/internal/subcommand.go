package internal

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCommand)
	RootCmd.AddCommand(startCommand)
	RootCmd.AddCommand(versionCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize config",
	Long:  "Initialize config",
	Run: func(cmd *cobra.Command, args []string) {
		InitStart()
	},
}

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "Start indexer",
	Long:  "Start indexer",
	Run: func(cmd *cobra.Command, args []string) {
		SendTx()
	},
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Long:  "Print the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		getVersion()
	},
}
