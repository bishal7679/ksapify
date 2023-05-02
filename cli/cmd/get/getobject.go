/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/spf13/cobra"
)

// getobjectCmd represents the getobject command
var GetobjectCmd = &cobra.Command{
	Use:     "get",
	Short:   "Use to get any kubernetes object",
	Aliases: []string{"get"},
	Long: `It is used to get any list of k8s object/resources. For example:

ksapify get ["pod", "deployment", "replicaset", "service", "..."]`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// cmd.RootCmd.AddCommand(GetobjectCmd)
}
