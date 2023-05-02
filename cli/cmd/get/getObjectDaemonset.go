/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectDaemonsetCmd represents the getObjectDaemonset command
var getObjectDaemonsetCmd = &cobra.Command{
	Use:   "daemonset",
	Short: "Use to get daemonset list inside the Cluster",
	Long: `It is used to get the list of daemonset inside the Cluster. For example:

ksapify get daemonset <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.Daemonsets(GetClusterns, Output, Wide)

	},
}

var (
	Output string
	Wide   bool
)

func init() {
	getObjectDaemonsetCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	getObjectDaemonsetCmd.Flags().StringVarP(&Output, "output", "o", "", "Output type [json/yaml]")
	getObjectDaemonsetCmd.Flags().BoolVarP(&Wide, "wide", "w", false, "Wider output [true/false]")
	GetobjectCmd.AddCommand(getObjectDaemonsetCmd)
}
