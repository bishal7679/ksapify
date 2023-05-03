package api

import (
	"context"
	"strconv"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// This func will create any pod
func CreatePod(Clusterns string, Podname string, Podimage string, Ports string) {
	clientset := Kconfig
	var hostPort, contPort int64
	ctx := context.Background()

	if Podname == "" || Podimage == "" || Ports == "" {
		logging.Err("🚨 podname and podimage and Ports should be provided")
		return
	}
	ports := strings.Split(Ports, ":")
	ports[0] = strings.Trim(ports[0], " ")
	ports[1] = strings.Trim(ports[1], " ")

	hostPort, _ = strconv.ParseInt(ports[0], 10, 32)
	contPort, _ = strconv.ParseInt(ports[1], 10, 32)
	ns, _ = CurrentNs(Clusterns)

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      Podname,
			Namespace: ns,
			Labels: map[string]string{
				"app": Podname,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  Podname + "-container",
					Image: Podimage,
					Ports: []v1.ContainerPort{
						{
							Name:          "http",
							Protocol:      v1.ProtocolTCP,
							HostPort:      int32(hostPort),
							ContainerPort: int32(contPort),
						},
					},
				},
			},
		},
	}

	podcreate, err := clientset.CoreV1().Pods(ns).Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		logging.Err("Error creating pod 😢")
		logging.Err(err.Error())
	} else {
		logging.Print("pod/" + podcreate.GetObjectMeta().GetName() + " created 🎉")
	}
}

// This func will delete any pod
func DeletePod(Clusterns string, PodName string) {
	clientset := Kconfig
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	err := clientset.CoreV1().Pods(ns).Delete(ctx, PodName, metav1.DeleteOptions{})
	if err != nil {
		logging.Err("Error deleting pod 😢")
		return
	}
	logging.Print("pod " + "\"" + PodName + "\"" + " deleted 🎉")
}

// This func will create any service
func CreateService(Clusterns string, Podname string, Servicename string, Servicetype string, Ports string, Nodeport int32) {
	clientset := Kconfig
	ctx := context.Background()
	var Sv string
	ports := strings.Split(Ports, ":")
	ports[0] = strings.Trim(ports[0], " ")
	ports[1] = strings.Trim(ports[1], " ")

	servicePort, _ := strconv.ParseInt(ports[0], 10, 32)
	targetPort, _ := strconv.ParseInt(ports[1], 10, 32)

	if Podname == "" || Servicename == "" || Ports == "" {
		logging.Err("🚨 podname, servicename, ports should be provided")
		return
	}
	ns, _ = CurrentNs(Clusterns)

	switch strings.ToLower(Servicetype) {
	case "nodeport":
		Sv = "NodePort"

	case "loadbalancer":
		Sv = "LoadBalancer"
	default:
		Sv = "ClusterIP"

	}

	// defining service manifest
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      Servicename,
			Namespace: ns,
			Labels: map[string]string{
				"app": Servicename + "-service",
			},
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceType(Sv),
			Ports: []v1.ServicePort{
				{
					Name:     Servicename + "-port",
					Protocol: v1.ProtocolTCP,
					// AppProtocol: v1.URISchemeHTTP,
					Port: int32(servicePort),
					TargetPort: intstr.IntOrString{
						IntVal: int32(targetPort),
					},
				},
			},
			Selector: map[string]string{
				"app": Podname,
			},
		},
	}

	if service.Spec.Type == "NodePort" {

		service.Spec.Ports = append(service.Spec.Ports, v1.ServicePort{NodePort: Nodeport})
	}
	servicecreate, err := clientset.CoreV1().Services(ns).Create(ctx, service, metav1.CreateOptions{})
	if err != nil {
		logging.Err("Error creating service 😢")
		logging.Err(err.Error())
	} else {
		logging.Print("service/" + servicecreate.GetObjectMeta().GetName() + " created 🎉")
	}

}

// This func will delete any service
func DeleteService(Clusterns string, ServiceName string) {
	clientset := Kconfig
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	err := clientset.CoreV1().Services(ns).Delete(ctx, ServiceName, metav1.DeleteOptions{})
	if err != nil {
		logging.Err("Error deleting service 😢")
		logging.Err(err.Error())
	}
	logging.Print("service " + "\"" + ServiceName + "\"" + " deleted 🎉")
}

// This func will create any deployment

func CreateDeployment(Clusterns string, Deploymentname string, Podimage string, Containerport int32) {
	clientset := Kconfig
	ctx := context.Background()

	if Deploymentname == "" || Podimage == "" {
		logging.Err("🚨 deploymentname and podimage should be provided")
		return
	}
	ns, _ = CurrentNs(Clusterns)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: Deploymentname,
			Labels: map[string]string{
				"app": Deploymentname,
			},
		},
		Spec: appsv1.DeploymentSpec{
			// Replicas: 2,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": Deploymentname,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": Deploymentname,
					},
				},
				Spec: v1.PodSpec{
					RestartPolicy: v1.RestartPolicyAlways,
					Containers: []v1.Container{
						{
							Name:            Deploymentname + "-container",
							Image:           Podimage,
							ImagePullPolicy: v1.PullIfNotPresent,
							Ports: []v1.ContainerPort{
								{
									Name:          "http",
									Protocol:      v1.ProtocolTCP,
									ContainerPort: Containerport,
								},
							},
						},
					},
				},
			},
		},
	}
	deploymentcreate, err := clientset.AppsV1().Deployments(ns).Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		logging.Err("Error creating deployment 😢")
		logging.Err(err.Error())
	} else {
		logging.Print("deployment/" + deploymentcreate.GetObjectMeta().GetName() + " created 🎉")
	}

}

// This function Deletes the Deployments
func DeleteDeployment(Clusterns string, DeploymentName string) {
	clientset := Kconfig
	ctx := context.Background()
	ns, _ = CurrentNs(Clusterns)
	err := clientset.AppsV1().Deployments(ns).Delete(ctx, DeploymentName, metav1.DeleteOptions{})
	if err != nil {
		logging.Err("Error deleting deployment 😢")
		logging.Err(err.Error())
	}
	logging.Print("deployment " + "\"" + DeploymentName + "\"" + " deleted 🎉")
}

// This func will create namespace
func CreateNamespace(Namespacename string) {
	clientset := Kconfig
	ctx := context.Background()
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: Namespacename,
			Labels: map[string]string{
				"name": Namespacename,
			},
		},
	}
	namespacecreate, err := clientset.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
	if err != nil {
		logging.Err("Error creating namespace 😢")
		logging.Err(err.Error())
	} else {
		logging.Print("namespace/" + namespacecreate.GetObjectMeta().GetName() + " created 🎉")
	}
}

// This func will delete namespace
func DeleteNamespace(Namespacename string) {
	clientset := Kconfig
	ctx := context.Background()
	err := clientset.CoreV1().Namespaces().Delete(ctx, Namespacename, metav1.DeleteOptions{})
	if err != nil {
		logging.Err("Error deleting namespace 😢")
		logging.Err(err.Error())
	}
	logging.Print("namespace " + "\"" + Namespacename + "\"" + " deleted 🎉")
}
