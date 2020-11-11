package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "homectl",
		Short: "homectl is a command line utility for your smart home",
		Long:  `A fast easy to use command line utility for use with your smart home`,
	}

	cfgFile string
)

func init() {
	cobra.OnInitialize()

         // Add flags to the root command
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "~/.config/homectl.yaml", "config file (default is $HOME/.config/homectl.yaml)")
	rootCmd.MarkPersistentFlagRequired("config")
}

func main() {
	_ = rootCmd.Execute()
}
