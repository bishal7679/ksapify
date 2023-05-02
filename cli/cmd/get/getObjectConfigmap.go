/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectConfigmapCmd represents the getObjectConfigmap command
var getObjectConfigmapCmd = &cobra.Command{
	Use:   "configmap",
	Short: "Use to get configmap list inside the Cluster",
	Long: `It is used to get the list of configmap inside the Cluster. For example:

ksapify get configmap <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.Configmaps(GetClusterns, Output, Wide)
	},
}

func init() {
	getObjectConfigmapCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	getObjectConfigmapCmd.Flags().StringVarP(&Output, "output", "o", "", "Output type [json/yaml]")
	getObjectConfigmapCmd.Flags().BoolVarP(&Wide, "wide", "w", false, "Wider output [true/false]")
	GetobjectCmd.AddCommand(getObjectConfigmapCmd)

}
