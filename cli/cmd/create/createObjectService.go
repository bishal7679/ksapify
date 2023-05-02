/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"strings"

	"github.com/bishal7679/ksapify/apis/api"
	"github.com/spf13/cobra"
)

// createObjectServiceCmd represents the createObjectService command
var createObjectServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Use to create a service",
	Long: `It is used to create a service inside the cluster. For example:

ksapify create service --namespace default --podname nginx-pod --name nginx-service --type loadbalancer --ports 8081:80   **(you can also use shorthand)**`,
	Run: func(cmd *cobra.Command, args []string) {
		api.OutsideClusterConfig()
		api.CreateService(GetClusterns, PodName, Servicename, Servicetype, Port, Nodeport)
	},
}

var (
	GetClusterns string
	PodName      string
)

var (
	Servicename string
	Servicetype string
	Port        string
	Nodeport    int32
)

func init() {
	createObjectServiceCmd.Flags().StringVarP(&GetClusterns, "namespace", "n", "", "Namespace name")
	createObjectServiceCmd.Flags().StringVarP(&PodName, "podname", "", "", "Pod name in which this service will be attached")
	createObjectServiceCmd.Flags().StringVarP(&Servicename, "name", "N", "", "Service name")
	createObjectServiceCmd.Flags().StringVarP(&Servicetype, "type", "t", "", "Service type")
	createObjectServiceCmd.Flags().StringVarP(&Port, "ports", "p", "", "Service port [servicePort:targetPort]")
	if strings.ToLower(Servicetype) == "nodeport" {
		createObjectServiceCmd.Flags().Int32VarP(&Nodeport, "nodeport", "", 34444, "Service nodeport")
		createObjectServiceCmd.MarkFlagRequired("nodeport")

	}
	CreateobjectCmd.AddCommand(createObjectServiceCmd)
	createObjectServiceCmd.MarkFlagRequired("podname")
	createObjectServiceCmd.MarkFlagRequired("name")
	createObjectServiceCmd.MarkFlagRequired("type")
	createObjectServiceCmd.MarkFlagRequired("ports")

}
