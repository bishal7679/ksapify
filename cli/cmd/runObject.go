/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// runObjectCmd represents the runObject command
var runObjectCmd = &cobra.Command{
	Use:   "run",
	Short: "Use to run any pod instantly",
	Long: `It is used to run any pod instantly inside the cluster. For example:

ksapify run --name nginx --image nginx:latest --ports 80:80  **(you can also use shorthand)**`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.CreatePod(GetClusterns, PodName, PodImage, Ports)
	},
}

var (
	GetClusterns string
	PodName      string
	PodImage     string
	Ports        string
)

func init() {
	runObjectCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	runObjectCmd.Flags().StringVar(&PodName, "name", "", "Name of pod")
	runObjectCmd.Flags().StringVarP(&PodImage, "image", "i", "", "Pod Image name")
	runObjectCmd.Flags().StringVarP(&Ports, "ports", "p", "", "Port [hostport:containerport]")
	RootCmd.AddCommand(runObjectCmd)
	runObjectCmd.MarkFlagRequired("name")
	runObjectCmd.MarkFlagRequired("image")
	runObjectCmd.MarkFlagRequired("ports")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runObjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runObjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
