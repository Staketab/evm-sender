package internal

import (
	"github.com/spf13/cobra"
	vars "github.com/staketab/evm-sender/cmd/evm-sender/internal/var"
	"sync"
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
	Short: "Start EVM Sender",
	Long:  "Start EVM Sender",
	Run: func(cmd *cobra.Command, args []string) {
		vars.InitLogger()
		logger := vars.GetLogger()
		var wg sync.WaitGroup
		done := make(chan bool)
		wg.Add(2)
		go func() {
			defer wg.Done()
			SendRangeTx(logger)
			done <- true
		}()
		//sendBackTicker := vars.SendBackTicker
		go func() {
			defer wg.Done()
			<-done
			//<-sendBackTicker.C
			SendBackTx(logger)
		}()
		wg.Wait()
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
