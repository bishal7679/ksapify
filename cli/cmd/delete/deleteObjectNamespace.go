/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// deleteObjectNamespaceCmd represents the deleteObjectNamespace command
var deleteObjectNamespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Use to delete a namespace",
	Long: `It is used to delete a namespace inside the Cluster. For example:

ksapify delete namespace <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.DeleteNamespace(Name)
	},
}

var (
	Name string
)

func init() {
	deleteObjectNamespaceCmd.Flags().StringVarP(&Name, "name", "n", "", "Namespace name")
	DeleteobjectCmd.AddCommand(deleteObjectNamespaceCmd)
	deleteObjectNamespaceCmd.MarkFlagRequired("name")

}
