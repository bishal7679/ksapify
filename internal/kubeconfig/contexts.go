package kubeconfig

import (
	"github.com/pkg/errors"

	"gopkg.in/yaml.v3"
)

// Use to get contexts node from Root node
func (k *Kubeconfig) contextsNode() (*yaml.Node, error) {
	contexts := valueOf(k.rootNode, "contexts")
	if contexts == nil {
		return nil, errors.New("\"contexts\" field is nil")
	} else if contexts.Kind != yaml.SequenceNode {
		return nil, errors.New("\"contexts\" is not a sequence node")
	}
	return contexts, nil
}

// Use to get all named context under contexts node from Root node
func (k *Kubeconfig) contextNode(name string) (*yaml.Node, error) {

	contexts, err := k.contextsNode()
	if err != nil {
		return nil, err
	}
	for _, contextNode := range contexts.Content {
		nameNode := valueOf(contextNode, "name")
		if nameNode.Kind == yaml.ScalarNode && nameNode.Value == name {
			return contextNode, nil
		}
	}
	return nil, errors.Errorf("context with name \"%s\" not found", name)
}

// this func is used to get any yaml node with all key-value
func valueOf(mapNode *yaml.Node, key string) *yaml.Node {
	if mapNode.Kind != yaml.MappingNode {
		return nil
	}
	for i, ch := range mapNode.Content {
		if i%2 == 0 && ch.Kind == yaml.ScalarNode && ch.Value == key {
			return mapNode.Content[i+1]
		}
	}
	return nil
}
