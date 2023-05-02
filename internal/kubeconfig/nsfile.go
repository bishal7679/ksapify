package kubeconfig

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bishal7679/ksapify/internal/util"
)

var defaultDir = filepath.Join(util.HomeDir(), ".kube", "ksapify")

type NamespaceFile struct {
	dir string
	ctx string
}

func NewNSFile(ctx string) NamespaceFile { return NamespaceFile{dir: defaultDir, ctx: ctx} }

func (f NamespaceFile) path() string {
	ctxname := f.ctx
	if isWindows() {
		ctxname = strings.ReplaceAll(ctxname, ":", "__")
	}
	return filepath.Join(f.dir, ctxname)
}

// Load reads the previous namespace setting, or returns empty if not exists.
func (f NamespaceFile) Load() (string, error) {
	b, err := os.ReadFile(f.path())
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return string(bytes.TrimSpace(b)), nil
}

// Save stores the previous namespace information in the file.
func (f NamespaceFile) Save(value string) error {
	d := filepath.Dir(f.path())
	if err := os.MkdirAll(d, 0755); err != nil {
		return err
	}
	return os.WriteFile(f.path(), []byte(value), 0644)
}

// isWindows determines if the process is running on windows OS.
func isWindows() bool {
	if os.Getenv("_FORCE_GOOS") == "windows" { // for testing
		return true
	}
	return runtime.GOOS == "windows"
}
