/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectReplicationcontrollerCmd represents the getObjectReplicationcontroller command
var getObjectReplicationcontrollerCmd = &cobra.Command{
	Use:   "replicationcontroller",
	Short: "Use to get replicationcontroller list inside the Cluster",
	Long: `It is used to get the list of replicationcontroller inside the Cluster. For example:

ksapify get replicationcontroller <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.Replicationcontrollers(GetClusterns, Output, Wide)

	},
}

func init() {
	getObjectReplicationcontrollerCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	getObjectReplicationcontrollerCmd.Flags().StringVarP(&Output, "output", "o", "", "Output type [json/yaml]")
	getObjectReplicationcontrollerCmd.Flags().BoolVarP(&Wide, "wide", "w", false, "Wider output [true/false]")
	GetobjectCmd.AddCommand(getObjectReplicationcontrollerCmd)

}
