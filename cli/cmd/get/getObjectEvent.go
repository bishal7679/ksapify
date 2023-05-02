/*
Copyright © 2023 NAME HERE bishalhnj127@gmail.com
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectEventCmd represents the getObjectEvent command
var getObjectEventCmd = &cobra.Command{
	Use:   "event",
	Short: "Use to get event list inside the Cluster",
	Long: `It is used to get the list of event inside the Cluster. For example:

ksapify get event <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.Events(GetClusterns, Output, Wide)

	},
}

func init() {
	getObjectEventCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	getObjectEventCmd.Flags().StringVarP(&Output, "output", "o", "", "Output type [json/yaml]")
	getObjectEventCmd.Flags().BoolVarP(&Wide, "wide", "w", false, "Wider output [true/false]")
	GetobjectCmd.AddCommand(getObjectEventCmd)
}
