/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// deleteObjectDeploymentCmd represents the deleteObjectDeployment command
var deleteObjectDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Use to delete any pod",
	Long: `It is used to delete any deployment inside the Cluster. For example:

ksapify delete deployment <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.DeleteDeployment(GetClusterns, Name)
	},
}

func init() {
	deleteObjectDeploymentCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	deleteObjectDeploymentCmd.Flags().StringVarP(&Name, "name", "N", "", "Deployment name")
	DeleteobjectCmd.AddCommand(deleteObjectDeploymentCmd)
	deleteObjectDeploymentCmd.MarkFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteObjectDeploymentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteObjectDeploymentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
