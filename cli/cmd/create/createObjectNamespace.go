/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// createObjectNamespaceCmd represents the createObjectNamespace command
var createObjectNamespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Use to create a namespace",
	Long: `It is used to create a namespace inside the cluster. For example:

ksapify create namespace --name test-namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.CreateNamespace(NamespaceName)
	},
}

var (
	NamespaceName string
)

func init() {
	createObjectNamespaceCmd.Flags().StringVarP(&NamespaceName, "name", "n", "", "Name of the namespace to be created")
	CreateobjectCmd.AddCommand(createObjectNamespaceCmd)
	createObjectNamespaceCmd.MarkFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createObjectNamespaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createObjectNamespaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
