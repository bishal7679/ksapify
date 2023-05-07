/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// deleteObjectServiceCmd represents the deleteObjectService command
var deleteObjectServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Use to delete a Service",
	Long: `It is used to delete a service inside the Cluster. For example:

ksapify delete service <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.DeleteService(GetClusterns, Name)
	},
}

func init() {
	deleteObjectServiceCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	deleteObjectServiceCmd.Flags().StringVarP(&Name, "name", "N", "", "Service name")
	DeleteobjectCmd.AddCommand(deleteObjectServiceCmd)
	deleteObjectServiceCmd.MarkFlagRequired("name")
}
