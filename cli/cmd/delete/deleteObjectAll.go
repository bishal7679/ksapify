/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// deleteObjectAllCmd represents the deleteObjectAll command
var deleteObjectAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Use to delete all k8s object",
	Long: `It is used to delete all objects of a namespace inside the Cluster. For example:

ksapify delete all <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.DeleteAllObject(GetClusterns)
	},
}

func init() {
	deleteObjectAllCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	DeleteobjectCmd.AddCommand(deleteObjectAllCmd)
}
