/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectServiceCmd represents the getObjectService command
var getObjectServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Use to get service list inside the Cluster",
	Long: `It is used to get the list of services inside the Cluster. For example:

ksapify get service --namespace default  **(you can also use shorthand)**`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.Services(GetClusterns, Output, Wide)
	},
}

func init() {
	getObjectServiceCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	getObjectServiceCmd.Flags().StringVarP(&Output, "output", "o", "", "Output type [json/yaml]")
	getObjectServiceCmd.Flags().BoolVarP(&Wide, "wide", "w", false, "Wider output [true/false]")
	GetobjectCmd.AddCommand(getObjectServiceCmd)

}
