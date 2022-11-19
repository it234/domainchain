package main

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"domainchain/app/browser"
	"domainchain/app/chain"
	"domainchain/app/wallet"

	cmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	"github.com/tendermint/tendermint/cmd/tendermint/commands/debug"
	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/node"
)

func main() {
	serverChain := chain.Run()
	defer serverChain.Stop()

	os.Args = append(os.Args, "run", "--proxy-app", "tcp://127.0.0.1:9080", "--home", "./confdata/tendermint-home")
	go func() {
		rootCmd := cmd.RootCmd
		rootCmd.AddCommand(
			cmd.GenValidatorCmd,
			cmd.ReIndexEventCmd,
			cmd.InitFilesCmd,
			cmd.ProbeUpnpCmd,
			cmd.LightCmd,
			cmd.ReplayCmd,
			cmd.ReplayConsoleCmd,
			cmd.ResetAllCmd,
			cmd.ResetPrivValidatorCmd,
			cmd.ResetStateCmd,
			cmd.ShowValidatorCmd,
			cmd.TestnetFilesCmd,
			cmd.ShowNodeIDCmd,
			cmd.GenNodeKeyCmd,
			cmd.VersionCmd,
			cmd.InspectCmd,
			cmd.RollbackStateCmd,
			cmd.MakeKeyMigrateCommand(),
			cmd.MakeCompactDBCommand(),
			debug.DebugCmd,
			cli.NewCompletionCmd(rootCmd, true),
		)
		nodeFunc := node.NewDefault
		rootCmd.AddCommand(cmd.NewRunNodeCmd(nodeFunc))

		cmd := cli.PrepareBaseCmd(rootCmd, "TM", os.ExpandEnv(filepath.Join("$HOME", config.DefaultTendermintDir)))
		if err := cmd.Execute(); err != nil {
			panic(err)
		}
	}()

	serverBrowser := browser.Run()
	defer serverBrowser.Close()

	serverWallet := wallet.Run()
	defer serverWallet.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	os.Exit(0)
}
