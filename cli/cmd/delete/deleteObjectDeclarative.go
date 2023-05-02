/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// deleteObjectDeclarativeCmd represents the deleteObjectDeclarative command
var DeleteObjectDeclarativeCmd = &cobra.Command{
	Use:   "decl",
	Short: "Use to delete all applied k8s object",
	Long: `It is used to delete declaratively created object inside the cluster. For example:

ksapify delete -f test.yaml  **(you can also use filename with full path location)**`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Declarative(Filename, true)
	},
}

var (
	Filename string
)

func init() {
	DeleteObjectDeclarativeCmd.Flags().StringVarP(&Filename, "file", "f", "", "Put Filename with proper location")
	DeleteobjectCmd.AddCommand(DeleteObjectDeclarativeCmd)
	DeleteObjectDeclarativeCmd.MarkFlagRequired("file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteObjectDeclarativeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteObjectDeclarativeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
