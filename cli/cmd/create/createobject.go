/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"github.com/spf13/cobra"
)

// createobjectCmd represents the createobject command
var CreateobjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Use to create any kubernetes object",
	Long: `It is used to create any k8s object/resources. For example:

ksapify create ["pod", "deployment", "replicaset", "service", "..."]`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
