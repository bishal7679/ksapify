/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	Log "github.com/bishal7679/ksapify/internal/logger"
	"github.com/spf13/cobra"
)

// getObjectPodNameCmd represents the getObjectPodName command
var getObjectPodCmd = &cobra.Command{
	Use:        "pod",
	Short:      "Use to get pod list inside the Cluster",
	SuggestFor: []string{"--namespace", "--contdet"},
	Long: `It is used to get the list of k8s pod inside the Cluster. For example:

ksapify get pod --namespace default --contdet=true   (you can also use shorthand)`,
	Run: func(cmd *cobra.Command, args []string) {

		api.OutsideClusterConfig()
		if IsContainerDetails {
			result := api.PodDetails(GetClusterns, IsContainerDetails)
			if result != "" {
				logging.Info(result, "")

			}
			return

		} else {
			err := api.Pods(GetClusterns)
			if err != nil {
				logging.Err(err.Error())
			}

		}

	},
}

var (
	IsContainerDetails bool
	GetClusterns       string
	logging            Log.Logger
)

func init() {
	getObjectPodCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	// getObjectPodNameCmd.Flags().String("owide", "", "Pod list with details info")
	getObjectPodCmd.Flags().BoolVarP(&IsContainerDetails, "contdet", "c", false, "Container details")
	GetobjectCmd.AddCommand(getObjectPodCmd)
	// getObjectPodNameCmd.MarkFlagRequired("namespace")

}
