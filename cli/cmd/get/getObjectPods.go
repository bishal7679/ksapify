/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectPodsNameCmd represents the getObjectPodsName command
var getObjectPodsCmd = &cobra.Command{
	Use:        "pods",
	Short:      "Use to get pods list inside the Cluster",
	SuggestFor: []string{"--namespace", "--contdet"},
	Long: `It is used to get the list of k8s pods inside the Cluster. For example:

ksapify get pods --namespace default --contdet=true   **(you can also use shorthand)**`,
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

func init() {
	getObjectPodsCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	// getObjectPodNameCmd.Flags().String("owide", "", "Pod list with details info")
	getObjectPodsCmd.Flags().BoolVarP(&IsContainerDetails, "contdet", "c", false, "Container details")
	GetobjectCmd.AddCommand(getObjectPodsCmd)
	// getObjectPodsNameCmd.MarkFlagRequired("namespace")
}
