/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/bishal7679/ksapify/cli/cmd/apply"
	"github.com/bishal7679/ksapify/cli/cmd/create"
	"github.com/bishal7679/ksapify/cli/cmd/delete"
	"github.com/bishal7679/ksapify/cli/cmd/get"
	"github.com/bishal7679/ksapify/cli/cmd/switching"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ksapify",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ksapify.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.AddCommand(get.GetobjectCmd)
	RootCmd.AddCommand(create.CreateobjectCmd)
	RootCmd.AddCommand(delete.DeleteobjectCmd)
	RootCmd.AddCommand(delete.DeleteObjectDeclarativeCmd)
	RootCmd.AddCommand(apply.DeclarativeObjectCmd)
	RootCmd.AddCommand(switching.SwitchNSCmd)

}
