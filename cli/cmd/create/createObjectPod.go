/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// createObjectPodCmd represents the createObjectPod command
var createObjectPodCmd = &cobra.Command{
	Use:   "pod",
	Short: "Use to create a pod",
	Long: `It is used to create a pod inside the cluster. For example:

ksapify create pod --namespace default --name nginx-pod --image nginx:latest --ports 80:80   **(you can also use shorthand)**`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.CreatePod(GetClusterns, PodName, PodImage, Ports)
	},
}

var (
	PodImage string
	Ports    string
)

func init() {
	createObjectPodCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	createObjectPodCmd.Flags().StringVarP(&PodName, "name", "N", "", "Pod name")
	createObjectPodCmd.Flags().StringVarP(&PodImage, "image", "i", "", "Image name")
	createObjectPodCmd.Flags().StringVarP(&Ports, "ports", "p", "", "Port [hostport:containerport]")
	CreateobjectCmd.AddCommand(createObjectPodCmd)
	createObjectPodCmd.MarkFlagRequired("name")
	// createObjectPodCmd.MarkFlagRequired("namespace")
	createObjectPodCmd.MarkFlagRequired("ports")
	createObjectPodCmd.MarkFlagRequired("image")

}
