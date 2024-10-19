package main

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/stelmanjones/termtools/prompt"
)

var rootCmd = &cobra.Command{
	Use:   "neutron",
	Short: "Easy way to build apps with a webview window.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new project.",
	Run: func(cmd *cobra.Command, args []string) {
		n, err := prompt.Ask("Project name: ", false)
		if err != nil {
			logger.Error(err.Error())
		}
		wd, _ := os.Getwd()

		os.MkdirAll(path.Join(wd, n), 0777)

		if err != nil {
			logger.Error(err.Error())
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
