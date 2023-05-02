/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// getObjectNamespaceCmd represents the getObjectNamespace command
var getObjectNamespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Use to get namespace list inside the cluster",
	Long: `It is used to get the list of namespace inside the cluster. For example:

ksapify get namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		result := api.Namespaces()
		logging.Print(result)
	},
}

func init() {
	GetobjectCmd.AddCommand(getObjectNamespaceCmd)
}
