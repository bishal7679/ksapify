/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package switching

import (
	"io"

	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// switchNamespaceCmd represents the switchNamespace command
var SwitchNamespaceCmd = &cobra.Command{
	Use:   "ns",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.Run(Writer, Namespace)
	},
}

var (
	Namespace string
	Writer    io.Writer
)

func init() {
	SwitchNamespaceCmd.Flags().StringVarP(&Namespace, "name", "n", "", "Namespace name")
	SwitchNamespaceCmd.MarkFlagRequired("name")
	SwitchNSCmd.AddCommand(SwitchNamespaceCmd)
}
