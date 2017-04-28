package cmd

import (
	"github.com/spf13/cobra"
	"bjdaos_tool/pkg/cli"
)

var RootCmd = &cobra.Command{
	Use: "heart_data",
	Short: "bjdaos_tool App cmd",
	Long: `bjdaos_tool is a tool to manage golang servic`,
	Example:`
		bjdaos_tool hd
	`,
}

func init(){
	RootCmd.AddCommand(cli.HeartDataCmd("hd"))
}