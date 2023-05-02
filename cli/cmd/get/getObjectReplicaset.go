/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectReplicasetCmd represents the getObjectReplicaset command
var getObjectReplicasetCmd = &cobra.Command{
	Use:   "replicaset",
	Short: "Use to get replicaset list inside the Cluster",
	Long: `It is used to get the list of replicaset inside the Cluster. For example:

ksapify get replicaset <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.Replicasets(GetClusterns, Output, Wide)
	},
}

func init() {
	getObjectReplicasetCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	getObjectReplicasetCmd.Flags().StringVarP(&Output, "output", "o", "", "Output type [json/yaml]")
	getObjectReplicasetCmd.Flags().BoolVarP(&Wide, "wide", "w", false, "Wider output [true/false]")
	GetobjectCmd.AddCommand(getObjectReplicasetCmd)

}
