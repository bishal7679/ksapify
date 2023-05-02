/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package apply

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// declarativeObjectCmd represents the declarativeObject command
var DeclarativeObjectCmd = &cobra.Command{
	Use:   "apply",
	Short: "Use to apply k8s object configuration",
	Long: `It is used to apply k8s object configuration through a file inside the cluster. For example:

ksapify apply -f test.yaml **(you can also use filename with full path location)**`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Declarative(Filename, false)
	},
}

var (
	Filename string
)

func init() {
	DeclarativeObjectCmd.Flags().StringVarP(&Filename, "file", "f", "", "Put Filename with proper location")
	DeclarativeObjectCmd.MarkFlagRequired("file")
}
