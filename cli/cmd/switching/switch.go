/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package switching

import (
	"github.com/spf13/cobra"
)

// switchNSCmd represents the switchNS command
var SwitchNSCmd = &cobra.Command{
	Use:   "switch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
