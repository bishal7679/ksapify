/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// deleteObjectPodCmd represents the deleteObjectPod command
var deleteObjectPodCmd = &cobra.Command{
	Use:   "pod",
	Short: "Use to delete any pod",
	Long: `It is used to delete any pod inside the Cluster. For example:

ksapify delete pod <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.DeletePod(GetClusterns, Name)
	},
}
var (
	GetClusterns string
)

func init() {
	deleteObjectPodCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	deleteObjectPodCmd.Flags().StringVarP(&Name, "name", "N", "", "Pod name")
	DeleteobjectCmd.AddCommand(deleteObjectPodCmd)
	deleteObjectPodCmd.MarkFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteObjectPodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteObjectPodCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
