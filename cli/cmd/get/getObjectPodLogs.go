/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectPodLogsCmd represents the getObjectPodLogs command
var getObjectPodLogsCmd = &cobra.Command{
	Use:   "podlogs",
	Short: "Use to get pod logs inside the Cluster",
	Long: `It is used to get the logs of any pod inside the Cluster. For example:

ksapify get podlogs --podname nginx --namespace default   **(you can also use shorthand)**`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.PodLogs(getClusterns, PodName)
	},
}

var (
	getClusterns string
	PodName      string
)

func init() {
	getObjectPodLogsCmd.Flags().StringVarP(&getClusterns, "namespace", "n", "", "Namespace name")
	getObjectPodLogsCmd.Flags().StringVarP(&PodName, "podname", "p", "", "Pod name")
	getObjectPodLogsCmd.MarkFlagRequired("podname")
	getObjectPodLogsCmd.MarkFlagRequired("namespace")
	GetobjectCmd.AddCommand(getObjectPodLogsCmd)

}
