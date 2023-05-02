package api

import (
	"context"
	"io"
	"os"

	"github.com/bishal7679/ksapify/internal/kubeconfig"
	"github.com/pkg/errors"
	errors2 "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Run(stderr io.Writer, namespace string) {
	kc := new(kubeconfig.Kubeconfig).WithLoader(kubeconfig.DefaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		logging.Err(err.Error())
		return
	}

	toNS, err := switchNamespace(kc, namespace)
	if err != nil {
		logging.Err(err.Error())
		return
	}
	logging.Normal("✔ Active namespace is " + "\"" + toNS + "\"")
}

func switchNamespace(kc *kubeconfig.Kubeconfig, ns string) (string, error) {
	ctx := kc.GetCurrentContext()
	if ctx == "" {
		return "", errors.New("current-context is not set")
	}
	curNS, err := kc.NamespaceOfContext(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to get current namespace")
	}

	f := kubeconfig.NewNSFile(ctx)
	prev, err := f.Load()
	if err != nil {
		return "", errors.Wrap(err, "failed to load previous namespace from file")
	}

	if ns == "-" {
		if prev == "" {
			return "", errors.Errorf("No previous namespace found for current context (%s)", ctx)
		}
		ns = prev
	}

	ok, err := namespaceExists(kc, ns)
	if err != nil {
		return "", errors.Wrap(err, "failed to query if namespace exists (is cluster accessible?)")
	}
	if !ok {
		return "", errors.Errorf("no namespace exists with name \"%s\"", ns)
	}

	if err := kc.SetNamespace(ctx, ns); err != nil {
		return "", errors.Wrapf(err, "failed to change to namespace \"%s\"", ns)
	}
	if err := kc.Save(); err != nil {
		return "", errors.Wrap(err, "failed to save kubeconfig file")
	}
	if curNS != ns {
		if err := f.Save(curNS); err != nil {
			return "", errors.Wrap(err, "failed to save the previous namespace to file")
		}
	}
	return ns, nil
}

func namespaceExists(kc *kubeconfig.Kubeconfig, ns string) (bool, error) {
	// for tests
	if os.Getenv("_MOCK_NAMESPACES") != "" {
		return ns == "ns1" || ns == "ns2", nil
	}

	clientset, err := newKubernetesClientSet(kc)
	if err != nil {
		return false, errors.Wrap(err, "failed to initialize k8s REST client")
	}

	namespace, err := clientset.CoreV1().Namespaces().Get(context.Background(), ns, metav1.GetOptions{})
	if errors2.IsNotFound(err) {
		return false, nil
	}
	return namespace != nil, errors.Wrapf(err, "failed to query "+
		"namespace %q from k8s API", ns)
}

func newKubernetesClientSet(kc *kubeconfig.Kubeconfig) (*kubernetes.Clientset, error) {
	b, err := kc.Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert in-memory kubeconfig to yaml")
	}
	cfg, err := clientcmd.RESTConfigFromKubeConfig(b)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize config")
	}
	return kubernetes.NewForConfig(cfg)
}
