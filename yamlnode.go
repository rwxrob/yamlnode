package yamlnode

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

var Kinds = map[yaml.Kind]string{
	yaml.DocumentNode: `Document`,
	yaml.SequenceNode: `Sequence`,
	yaml.MappingNode:  `Mapping`,
	yaml.ScalarNode:   `Scalar`,
	yaml.AliasNode:    `Alias`,
}

// FromAsDoc returns a reference to a [gopkg.in/yaml.v3.Node] by first
// marshaling the argument into a YAML document string and then
// unmarshaling into a new Node and returning its reference along with
// the containing DocumentNode. When just the Node itself is wanted use
// [Cast] instead.
func FromAsDoc(this any) (*yaml.Node, error) {
	byt, err := yaml.Marshal(this)
	if err != nil {
		return nil, err
	}
	node := new(yaml.Node)
	if err := yaml.Unmarshal(byt, node); err != nil {
		return nil, err
	}
	return node, nil
}

// From returns a reference to a [gopkg.in/yaml.v3.Node] by first
// marshaling the argument into a YAML document string and then
// unmashalling into a new Node and returning its reference (Content[0]).
// The Node can then be transformed or queried by a [WalkFunc] with
// [Walk]. When the encapsulating document is also wanted use [FromAsDoc]
// instead.
func From(this any) (*yaml.Node, error) {
	node, err := FromAsDoc(this)
	if err != nil {
		return nil, err
	}
	return node.Content[0], nil
}

// ToString returns the Node marshaled into a YAML string including
// a nil argument as the string "null". If an error (unlikely) occurs
// while marshaling a "null" string is returned.
func ToString(node *yaml.Node) string {
	if node == nil {
		return "null"
	}
	byt, err := yaml.Marshal(node)
	if err != nil {
		return "null"
	}
	return string(byt)
}

// Print calls [ToString] and outputs it to [os.Stdout].
func Print(node *yaml.Node) { fmt.Print(ToString(node)) }
