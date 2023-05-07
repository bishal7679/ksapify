package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	cnfg "github.com/bishal7679/ksapify/internal/config"
	"github.com/bishal7679/ksapify/internal/kubeconfig"
	yml "github.com/ghodss/yaml"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	result string
	kc     *kubeconfig.Kubeconfig
	ns     string
)

// This function is used to get the all the list of pods name only
func Pods(Clusterns string) error {
	clientset := Kconfig
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	pods, err := clientset.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{})

	if err != nil {
		panic(err)
	}
	if len(pods.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
	} else {
		for i := 0; i < len(pods.Items); i++ {
			logging.Print(pods.Items[i].Name)
		}
	}
	return err
}

// This function is used to get the all the list of pods in details in the cluster with container details

func PodDetails(Clusterns string, ContainerDetails bool) string {

	clientset := Kconfig
	ctx := context.Background()

	var podInfo []cnfg.Pod
	var containerInfo []cnfg.Container

	ns, _ = CurrentNs(Clusterns)
	pods, err := clientset.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	if len(pods.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
	} else {
		for i := 0; i < len(pods.Items); i++ {
			podInfo = append(podInfo, cnfg.Pod{
				Name:            pods.Items[i].Name,
				Status:          string(pods.Items[i].Status.Phase),
				CreatedAt:       pods.Items[i].CreationTimestamp.String(),
				UniqueID:        string(pods.Items[i].GetUID()),
				NodeName:        string(pods.Items[i].Spec.NodeName),
				IP:              string(pods.Items[i].Status.PodIP),
				ContainersCount: len(pods.Items[i].Spec.Containers),
				Labels:          pods.Items[i].Labels,
			})
			if ContainerDetails {

				for j := 0; j < len(pods.Items[i].Spec.Containers); j++ {
					containerInfo = append(containerInfo, cnfg.Container{
						Name:            pods.Items[i].Spec.Containers[j].Name,
						Container:       j,
						Image:           pods.Items[i].Spec.Containers[j].Image,
						ImagePullPolicy: string(pods.Items[i].Spec.Containers[j].ImagePullPolicy),
						Port:            pods.Items[i].Spec.Containers[j].Ports,
					})
				}
			}
			podInfo[i].ContainersInfo = containerInfo
		}

		podsInfo_result, err := json.MarshalIndent(podInfo, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		result = string(podsInfo_result)
	}
	return result
}

func PodLogs(Clusterns string, PodName string) {
	clientset := Kconfig
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	pods, _ := clientset.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{})
	for i := 0; i < len(pods.Items); i++ {
		if pods.Items[i].Name == PodName {
			request := clientset.CoreV1().Pods(ns).GetLogs(PodName, &(v1.PodLogOptions{}))
			podLogs, err := request.Stream(ctx)
			if err != nil {
				logging.Err("error in opening stream *restclient.Request")
			}
			defer podLogs.Close()

			buf := new(bytes.Buffer)
			_, err = io.Copy(buf, podLogs)
			if err != nil {
				panic(err)
			}
			logging.Info(buf.String(), "")
			return
		}
	}
	logging.Err("Pod " + "\"" + PodName + "\"" + " not found 😢")

}

// This function is used to get the all the list of deployments in the cluster

func Deployments(Clusterns string, Output string, Wide bool) {
	clientset := Kconfig
	var deploymentlist []cnfg.Deployment
	var deploymentlistwide []cnfg.WideResult
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	deployments, err := clientset.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	if len(deployments.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
		return
	}
	if Wide || Output == "yaml" {
		for i := 0; i < len(deployments.Items); i++ {
			deploymentlistwide = append(deploymentlistwide, cnfg.WideResult{
				ApiVersion: v1.SchemeGroupVersion.Version,
				Items:      deployments.Items[i],
				Kind:       "List",
				Metadata:   deployments.ListMeta,
			})
		}

		result = OutputType(Output, deploymentlistwide)
		logging.Info(result, "")
		return
	} else {
		for i := 0; i < len(deployments.Items); i++ {
			deploymentlist = append(deploymentlist, cnfg.Deployment{
				Name:      deployments.Items[i].Name,
				Status:    string(deployments.Items[i].Status.Conditions[0].Type),
				CreatedAt: deployments.Items[i].CreationTimestamp.String(),
				UniqueID:  string(deployments.Items[i].UID),
				Labels:    deployments.Items[i].Labels,
			})
		}
		result = OutputType(Output, deploymentlist)
		logging.Info(result, "")
		return

	}

}

// This function is used to get the all the list of services in the cluster

func Services(Clusterns string, Output string, Wide bool) {

	clientset := Kconfig
	var servicelist []cnfg.Service
	var servicelistwide []cnfg.WideResult
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)

	services, err := clientset.CoreV1().Services(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	if len(services.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
		return
	}
	if Wide || Output == "yaml" {
		for i := 0; i < len(services.Items); i++ {
			servicelistwide = append(servicelistwide, cnfg.WideResult{
				ApiVersion: v1.SchemeGroupVersion.Version,
				Items:      services.Items[i],
				Kind:       "List",
				Metadata:   services.ListMeta,
			})
		}

		result = OutputType(Output, servicelistwide)
		logging.Info(result, "")
		return
	} else {
		for i := 0; i < len(services.Items); i++ {
			servicelist = append(servicelist, cnfg.Service{
				Name:      services.Items[i].Name,
				CreatedAt: services.Items[i].CreationTimestamp.String(),
				UniqueID:  string(services.Items[i].GetUID()),
				Labels:    services.Items[i].Labels,
				Spec:      services.Items[i].Spec,
			})
		}

		result = OutputType(Output, servicelist)
		logging.Info(result, "")
		return
	}
}

// This function is used to get the all the list of namespaces in the cluster

func Namespaces() string {

	clientset := Kconfig
	var namespacelist []cnfg.Namespace
	ctx := context.Background()
	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	} else {
		for i := 0; i < len(namespaces.Items); i++ {
			namespacelist = append(namespacelist, cnfg.Namespace{
				Name:      namespaces.Items[i].Name,
				CreatedAt: namespaces.Items[i].CreationTimestamp.String(),
				UniqueID:  string(namespaces.Items[i].UID),
			})
		}

		namespace_result, err := json.MarshalIndent(namespacelist, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		return string(namespace_result)

	}
}

// This function is used to get the all the list of configmap in the cluster

func Configmaps(Clusterns string, Output string, Wide bool) {

	clientset := Kconfig
	var configmaplist []cnfg.Configmap
	var configmaplistwide []cnfg.WideResult
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	configmaps, err := clientset.CoreV1().ConfigMaps(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	if len(configmaps.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
		return
	}
	if Wide || Output == "yaml" {
		for i := 0; i < len(configmaps.Items); i++ {
			configmaplistwide = append(configmaplistwide, cnfg.WideResult{
				ApiVersion: v1.SchemeGroupVersion.Version,
				Items:      configmaps.Items[i],
				Kind:       "List",
				Metadata:   configmaps.ListMeta,
			})
		}

		result = OutputType(Output, configmaplistwide)
		logging.Info(result, "")
		return
	} else {
		for i := 0; i < len(configmaps.Items); i++ {
			configmaplist = append(configmaplist, cnfg.Configmap{
				Name: configmaps.Items[i].Name,
			})
		}

		result = OutputType(Output, configmaplist)
		logging.Info(result, "")
		return
	}
}

// This function is used to get the all the list of daemonsets in the cluster

func Daemonsets(Clusterns string, Output string, Wide bool) {

	clientset := Kconfig
	var daemonsetlist []cnfg.Daemonset
	var daemonsetlistwide []cnfg.WideResult
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	daemonsets, err := clientset.AppsV1().DaemonSets(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	if len(daemonsets.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
		return
	}
	if Wide || Output == "yaml" {
		for i := 0; i < len(daemonsets.Items); i++ {
			daemonsetlistwide = append(daemonsetlistwide, cnfg.WideResult{
				ApiVersion: v1.SchemeGroupVersion.Version,
				Items:      daemonsets.Items[i],
				Kind:       "List",
				Metadata:   daemonsets.ListMeta,
			})
		}

		result = OutputType(Output, daemonsetlistwide)
		logging.Info(result, "")
		return
	} else {
		for i := 0; i < len(daemonsets.Items); i++ {
			daemonsetlist = append(daemonsetlist, cnfg.Daemonset{
				Name:      daemonsets.Items[i].Name,
				CreatedAt: daemonsets.Items[i].CreationTimestamp.String(),
				UniqueID:  string(daemonsets.Items[i].UID),
				Labels:    daemonsets.Items[i].Labels,
			})
		}

		result = OutputType(Output, daemonsetlist)
		logging.Info(result, "")
		return

	}
}

func OutputType(otype string, list interface{}) string {

	var res []byte
	output, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	if otype == "yaml" {
		res, _ = yml.JSONToYAML(output)
		return string(res)
	}
	if otype == "json" {

		return string(output)
	}
	return string(output)
}

// This function is used to get the all the list of events in the cluster

func Events(Clusterns string, Output string, Wide bool) {

	clientset := Kconfig
	var eventlist []cnfg.Event
	var eventlistwide []cnfg.WideResult
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	events, err := clientset.CoreV1().Events(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	if len(events.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
		return
	}
	if Wide || Output == "yaml" {
		for i := 0; i < len(events.Items); i++ {
			eventlistwide = append(eventlistwide, cnfg.WideResult{
				ApiVersion: v1.SchemeGroupVersion.Version,
				Items:      events.Items[i],
				Kind:       "List",
				Metadata:   events.ListMeta,
			})
		}

		result = OutputType(Output, eventlistwide)
		logging.Info(result, "")
		return
	} else {
		for i := 0; i < len(events.Items); i++ {
			eventlist = append(eventlist, cnfg.Event{
				Name:       events.Items[i].Name,
				Type:       events.Items[i].Type,
				ObjectName: events.Items[i].InvolvedObject.Name,
				CreatedAt:  events.Items[i].LastTimestamp.String(),
				UniqueID:   string(events.Items[i].UID),
			})
		}
		result = OutputType(Output, eventlist)
		logging.Info(result, "")
		return

	}
}

// This function is used to get the all the list of replicationController in the cluster

func Replicationcontrollers(Clusterns string, Output string, Wide bool) {

	clientset := Kconfig
	var replicationcontrollerlist []cnfg.Replicationcontroller
	var replicationcontrollerlistwide []cnfg.WideResult
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	replicationcontrollers, err := clientset.CoreV1().ReplicationControllers(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)

	}
	if len(replicationcontrollers.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
		return
	}
	if Wide || Output == "yaml" {
		for i := 0; i < len(replicationcontrollers.Items); i++ {
			replicationcontrollerlistwide = append(replicationcontrollerlistwide, cnfg.WideResult{
				ApiVersion: v1.SchemeGroupVersion.Version,
				Items:      replicationcontrollers.Items[i],
				Kind:       "List",
				Metadata:   replicationcontrollers.ListMeta,
			})
		}

		result = OutputType(Output, replicationcontrollerlistwide)
		logging.Info(result, "")
		return
	} else {
		for i := 0; i < len(replicationcontrollers.Items); i++ {
			replicationcontrollerlist = append(replicationcontrollerlist, cnfg.Replicationcontroller{
				Name:      replicationcontrollers.Items[i].Name,
				CreatedAt: replicationcontrollers.Items[i].CreationTimestamp.String(),
				UniqueID:  string(replicationcontrollers.Items[i].UID),
				Labels:    replicationcontrollers.Items[i].Labels,
			})
		}

		result = OutputType(Output, replicationcontrollerlist)
		logging.Info(result, "")
		return

	}
}

// This function is used to get the all the list of replicasets in the cluster

func Replicasets(Clusterns string, Output string, Wide bool) {

	clientset := Kconfig
	var replicasetlist []cnfg.Replicaset
	var replicasetlistwide []cnfg.WideResult
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	replicasets, err := clientset.AppsV1().ReplicaSets(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	if len(replicasets.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
		return
	}
	if Wide || Output == "yaml" {
		for i := 0; i < len(replicasets.Items); i++ {
			replicasetlistwide = append(replicasetlistwide, cnfg.WideResult{
				ApiVersion: v1.SchemeGroupVersion.Version,
				Items:      replicasets.Items[i],
				Kind:       "List",
				Metadata:   replicasets.ListMeta,
			})
		}

		result = OutputType(Output, replicasetlistwide)
		logging.Info(result, "")
		return
	} else {
		for i := 0; i < len(replicasets.Items); i++ {
			replicasetlist = append(replicasetlist, cnfg.Replicaset{
				Name:           replicasets.Items[i].Name,
				CreatedAt:      replicasets.Items[i].CreationTimestamp.String(),
				UniqueID:       string(replicasets.Items[i].UID),
				ReplicasetSpec: replicasets.Items[i].Spec,
				Labels:         replicasets.Items[i].Labels,
			})
		}

		result = OutputType(Output, replicasetlist)
		logging.Info(result, "")
		return

	}
}

// This function is used to get the list of all the secrets in the cluster

func Secrets(Clusterns string, Output string, Wide bool) {
	clientset := Kconfig
	var secretlist []cnfg.Secret
	var secretlistwide []cnfg.WideResult
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	secrets, err := clientset.CoreV1().Secrets(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	if len(secrets.Items) == 0 {
		logging.Print("No resources found in " + ns + " namespace")
		return
	}
	if Wide || Output == "yaml" {
		for i := 0; i < len(secrets.Items); i++ {
			secretlistwide = append(secretlistwide, cnfg.WideResult{
				ApiVersion: v1.SchemeGroupVersion.Version,
				Items:      secrets.Items[i],
				Kind:       "List",
				Metadata:   secrets.ListMeta,
			})
		}

		result = OutputType(Output, secretlistwide)
		logging.Info(result, "")
		return
	} else {
		for i := 0; i < len(secrets.Items); i++ {
			secretlist = append(secretlist, cnfg.Secret{
				Name:      secrets.Items[i].Name,
				CreatedAt: secrets.Items[i].CreationTimestamp.String(),
				UniqueID:  string(secrets.Items[i].UID),
			})
			scrttmp := make(map[string]string)
			for key, val := range secrets.Items[i].Data {
				scrttmp[key] = string(val)
			}
			secretlist[i].SecretMap = scrttmp
		}

		result = OutputType(Output, secretlist)
		logging.Info(result, "")
		return
	}

}

// This function is used to apply declarative resource config in the cluster
func Declarative(filename string, Delete bool) {

	file, err := os.ReadFile(filename)
	if err != nil {
		logging.Err(err.Error())
	}
	kubeconfig := os.Getenv("KUBECONFIG")
	config, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logging.Err(err.Error())
		return
	}

	dynamicClientset, err := dynamic.NewForConfig(config)
	if err != nil {
		logging.Err(err.Error())
		return
	}

	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(file), 100)
	for {
		var rawObj runtime.RawExtension
		if err = decoder.Decode(&rawObj); err != nil {
			break
		}

		obj, grpvrsnkind, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		if err != nil {
			logging.Err(err.Error())
			return
		}

		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			logging.Err(err.Error())
			return
		}

		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

		grpresources, err := restmapper.GetAPIGroupResources(clientset.Discovery())
		if err != nil {
			logging.Err(err.Error())
			return
		}

		mapper := restmapper.NewDiscoveryRESTMapper(grpresources)
		mapping, err := mapper.RESTMapping(grpvrsnkind.GroupKind(), grpvrsnkind.Version)
		if err != nil {
			logging.Err(err.Error())
			return
		}

		var dri dynamic.ResourceInterface

		ctx := context.Background()
		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
			if unstructuredObj.GetNamespace() == "" {
				unstructuredObj.SetNamespace("default")
			}
			dri = dynamicClientset.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
		} else {
			dri = dynamicClientset.Resource(mapping.Resource)
		}

		if !Delete {
			if created, err := dri.Create(ctx, unstructuredObj, metav1.CreateOptions{}); err != nil {
				logging.Err(err.Error())
				fmt.Println("ddd")
				return
			} else {
				objkind := created.GetKind()
				objname := created.GetName()
				logging.Print(objkind + "/" + objname + " created 🎉")
			}
		}
		if Delete {
			if err := dri.DeleteCollection(ctx, metav1.DeleteOptions{}, metav1.ListOptions{}); err != nil {
				logging.Err(err.Error())
				return
			} else {
				logging.Print(unstructuredObj.GetKind() + " " + "\"" + unstructuredObj.GetName() + "\"" + " deleted 🎉")
				// unstructuredObj.SetDeletionTimestamp(&metav1.Time{Time: time.Now()})
			}
		}
	}
	if err != io.EOF {
		logging.Err(err.Error())
		return
	}
}

func AllObject(Clusterns string) {
	OutsideClusterConfig()
	countobj := 0
	clientset := Kconfig
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	// pod names
	pods, _ := clientset.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{})

	if len(pods.Items) != 0 {
		logging.Normal("NAME")
		for i := 0; i < len(pods.Items); i++ {
			logging.Print("pod/" + pods.Items[i].Name)
		}
		fmt.Printf("\n")
		countobj += 1
	}

	// service names
	svc, _ := clientset.CoreV1().Services(ns).List(ctx, metav1.ListOptions{})

	if len(svc.Items) != 0 {
		logging.Normal("NAME")
		for i := 0; i < len(svc.Items); i++ {
			logging.Print("service/" + svc.Items[i].Name)
		}
		fmt.Printf("\n")
		countobj += 1
	}

	// daemonsetnames
	daemonsets, _ := clientset.AppsV1().DaemonSets(ns).List(ctx, metav1.ListOptions{})

	if len(daemonsets.Items) != 0 {
		logging.Normal("NAME")
		for i := 0; i < len(daemonsets.Items); i++ {
			logging.Print("daemonset.apps/" + daemonsets.Items[i].Name)
		}
		fmt.Printf("\n")
		countobj += 1
	}

	// deployment name
	deployments, _ := clientset.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{})

	if len(deployments.Items) != 0 {
		logging.Normal("NAME")
		for i := 0; i < len(deployments.Items); i++ {
			logging.Print("deployment.apps/" + deployments.Items[i].Name)
		}
		fmt.Printf("\n")
		countobj += 1
	}

	// replicaset name
	replicasets, _ := clientset.AppsV1().ReplicaSets(ns).List(ctx, metav1.ListOptions{})

	if len(replicasets.Items) != 0 {
		logging.Normal("NAME")
		for i := 0; i < len(replicasets.Items); i++ {
			logging.Print("replicaset.apps/" + replicasets.Items[i].Name)
		}
		countobj += 1
	}

	if countobj == 0 {
		logging.Print("No resources found in " + ns + " namespace")
	}

}

func CurrentNs(Clusterns string) (string, error) {
	var ns string

	kc = new(kubeconfig.Kubeconfig).WithLoader(kubeconfig.DefaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		logging.Err(err.Error())
	}
	ctx := kc.GetCurrentContext()
	if ctx == "" {
		return "", errors.New("current-context is not set")
	}
	curNS, err := kc.NamespaceOfContext(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to get current namespace")
	}
	if Clusterns == "" {
		ns = curNS
	} else {
		ns = Clusterns
	}
	return ns, nil
}
