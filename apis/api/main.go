package api

import (
	// "fmt"
	"os"
	"time"

	log "github.com/bishal7679/ksapify/internal/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	logging log.Logger
	Kconfig *kubernetes.Clientset
)

func Main() {
	logging.Info("ksapify is running successfully!🥳 🎉", "")

	// Uncomment this if you want to run outside the cluster
	OutsideClusterConfigFancy()
	// This will be used in case you have to run the code inside the cluster
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	logging.Err("error getting config inClusterconfig\n")
	// } else {
	// 	logging.Info("config built successfully!✅", "")
	// }

	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	logging.Err("error creating clientset 😢")
	// } else {
	// 	logging.Info("clientset created successfully!✅", "")
	// }

	// Kconfig = clientset

}

func OutsideClusterConfigFancy() {

	// You will have to export the KUBECONFIG variable to point to the kubeconfig file based on your OS in the terminal
	logging.Note("Export KUBECONFIG env variable is important!")
	kubeconfig := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		// fmt.Errorf("error %s building config from env variable\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			logging.Err("error getting config inClusterconfig 😢")
		}
	} else {
		time.Sleep(3 * time.Second)
		logging.Info("config built successfully ✔", "")
	}

	// creating the clientset to access all the resources
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logging.Err("error creating clientset 😢")
	} else {
		time.Sleep(4 * time.Second)
		logging.Info("clientset created successfully ✔", "")

	}

	Kconfig = clientset

}

func OutsideClusterConfig() {

	// You will have to export the KUBECONFIG variable to point to the kubeconfig file based on your OS in the terminal
	kubeconfig := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		// fmt.Errorf("error %s building config from env variable\n", err.Error())
		config, _ = rest.InClusterConfig()
	}

	// creating the clientset to access all the resources
	clientset, _ := kubernetes.NewForConfig(config)

	Kconfig = clientset

}
