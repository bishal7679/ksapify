package kubeconfig

import (
	"os"
	"path/filepath"

	"github.com/bishal7679/ksapify/internal/util"

	"github.com/pkg/errors"
)

var (
	DefaultLoader Loader = new(StandardKubeconfigLoader)
)

type StandardKubeconfigLoader struct{}

type kubeconfigFile struct{ *os.File }

func (*StandardKubeconfigLoader) Load() ([]ReadWriteResetCloser, error) {
	cgfPath, err := kubeconfigPath()
	if err != nil {
		return nil, errors.Wrap(err, "cannot determine kubeconfig path")
	}

	f, err := os.OpenFile(cgfPath, os.O_RDWR, 0)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.Wrap(err, "kubeconfig file not found")
		}
		return nil, errors.Wrap(err, "failed to open file")
	}

	return []ReadWriteResetCloser{ReadWriteResetCloser(&kubeconfigFile{f})}, nil
}

func (kf *kubeconfigFile) Reset() error {
	if err := kf.Truncate(0); err != nil {
		return errors.Wrap(err, "failed to truncate file")
	}
	_, err := kf.Seek(0, 0)
	return errors.Wrap(err, "failed to seek in file")
}

func kubeconfigPath() (string, error) {
	// KUBECONFIG env var
	if kcfg := os.Getenv("KUBECONFIG"); kcfg != "" {
		list := filepath.SplitList(kcfg)
		if len(list) > 1 {
			return "", errors.New("multiple files in KUBECONFIG are currently not supported")
		}
		return kcfg, nil
	}

	// set default path

	home := util.HomeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE environment variable not set")
	}
	return filepath.Join(home, ".kube", "config"), nil
}
