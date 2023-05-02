package kubeconfig

// GetCurrentContext returns "current-context" value from given kubeconfig root node yaml

func (k *Kubeconfig) GetCurrentContext() string {
	currentctx := valueOf(k.rootNode, "current-context")
	if currentctx == nil {
		return ""
	}
	return currentctx.Value
}

func (k *Kubeconfig) UnsetCurrentContext() error {
	curCtxValNode := valueOf(k.rootNode, "current-context")
	curCtxValNode.Value = ""
	return nil
}
