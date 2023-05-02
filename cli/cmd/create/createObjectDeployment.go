/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// createObjectDeploymentCmd represents the createObjectDeployment command
var createObjectDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Use to create a deployment",
	Long: `It is used to create a deployment inside the cluster. For example:

ksapify create deployment --namespace default --name nginx-deployment --image nginx:latest --contport 80   **(you can also use shorthand)**`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.CreateDeployment(GetClusterns, DeploymentName, PodImage, ContainerPort)
	},
}

var (
	ContainerPort  int32
	DeploymentName string
)

func init() {
	createObjectDeploymentCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	createObjectDeploymentCmd.Flags().StringVarP(&DeploymentName, "name", "N", "", "Deployment name")
	createObjectDeploymentCmd.Flags().StringVarP(&PodImage, "image", "i", "", "Image name")
	createObjectDeploymentCmd.Flags().Int32VarP(&ContainerPort, "contport", "p", 80, "")

	CreateobjectCmd.AddCommand(createObjectDeploymentCmd)
	createObjectDeploymentCmd.MarkFlagRequired("name")
	createObjectDeploymentCmd.MarkFlagRequired("image")
	createObjectDeploymentCmd.MarkFlagRequired("contport")
}
