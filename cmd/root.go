/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hexcraft",
	Short: "A CLI tool for generating hexagonal architecture for microservices in Go",
	Long: `Hexcraft is a command-line tool that helps you quickly generate the necessary files and structure
for creating a microservice using the hexagonal architecture pattern in Go.

The hexagonal architecture, also known as the ports and adapters architecture, provides a clean and
maintainable structure for building microservices. It separates the core business logic from the external
dependencies and infrastructure, making the code more modular and easier to test.

With Hexcraft, you can easily scaffold a new microservice project with the hexagonal architecture layout,
including the necessary packages, interfaces, and starter code. It saves you time and effort in setting up
the project structure manually.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hexcraft.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
