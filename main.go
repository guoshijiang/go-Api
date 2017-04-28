package main

import (
	"bjdaos_tool/pkg/cmd"
	"os"
)

func main(){
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}












