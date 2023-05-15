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
	Short: "Multi-Featured Light Kubernetes command-line tool",
	Long: `A Multi-Featured Light Kubernetes command-line tool which can interact with k8s api-server and perform operations`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.AddCommand(get.GetobjectCmd)
	RootCmd.AddCommand(create.CreateobjectCmd)
	RootCmd.AddCommand(delete.DeleteobjectCmd)
	RootCmd.AddCommand(delete.DeleteObjectDeclarativeCmd)
	RootCmd.AddCommand(apply.DeclarativeObjectCmd)
	RootCmd.AddCommand(switching.SwitchNSCmd)

}
