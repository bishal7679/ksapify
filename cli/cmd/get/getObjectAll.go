/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectAllCmd represents the getObjectAll command
var getObjectAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Use to get all k8s object",
	Long: `It is used to get all k8s object inside the cluster. For example:

ksapify get all --namespace default`,
	Run: func(cmd *cobra.Command, args []string) {
		api.AllObject(GetClusterns)
	},
}

func init() {
	getObjectAllCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	GetobjectCmd.AddCommand(getObjectAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getObjectAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getObjectAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
