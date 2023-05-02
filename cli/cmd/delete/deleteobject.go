/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"github.com/spf13/cobra"
)

// deleteobjectCmd represents the deleteobject command
var DeleteobjectCmd = &cobra.Command{
	Use:   "delete",
	Short: "Use to delete any kubernetes object",
	Long: `It is used to delete any object inside the cluster. For example:

ksapify delete <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
