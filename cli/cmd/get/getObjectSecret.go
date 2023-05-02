/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectSecretCmd represents the getObjectSecret command
var getObjectSecretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Use to get secret list inside the Cluster",
	Long: `It is used to get the list of secret inside the Cluster. For example:

ksapify get secret <use available flags>`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.Secrets(GetClusterns, Output, Wide)

	},
}

func init() {
	getObjectSecretCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	getObjectSecretCmd.Flags().StringVarP(&Output, "output", "o", "", "Output type [json/yaml]")
	getObjectSecretCmd.Flags().BoolVarP(&Wide, "wide", "w", false, "Wider output [true/false]")
	GetobjectCmd.AddCommand(getObjectSecretCmd)
}
