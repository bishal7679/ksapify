/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectDeploymentNameCmd represents the getObjectDeploymentName command
var getObjectDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Use to get deployment list inside the Cluster",
	Long: `It is used to get the list of deployment inside the Cluster. For example:

ksapify get deployment <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.Deployments(Clusterns, Output, Wide)
	},
}

var (
	Clusterns string
)

func init() {
	getObjectDeploymentCmd.Flags().StringVarP(&Clusterns, "namespace", "n", "", "Namespace name")
	getObjectDeploymentCmd.Flags().StringVarP(&Output, "output", "o", "", "Output type [json/yaml]")
	getObjectDeploymentCmd.Flags().BoolVarP(&Wide, "wide", "w", false, "Wider output [true/false]")
	GetobjectCmd.AddCommand(getObjectDeploymentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getObjectDeploymentNameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getObjectDeploymentNameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
