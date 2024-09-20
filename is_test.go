package yamlnode_test

import (
	"fmt"

	"github.com/rwxrob/yamlnode"
	"gopkg.in/yaml.v3"
)

func ExampleIsBool() {
	values := []any{
		true,
		false,
		`true`,
		`True`,
		`Yes`,
		`yes`,
	}
	for _, val := range values {
		node, _ := yamlnode.From(val)
		fmt.Println(yamlnode.IsBool(node))
	}
	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
}

func ExampleIsBool_from_Node() {
	node := new(yaml.Node)
	node.Kind = yaml.ScalarNode
	node.Tag = `!!bool`
	node.Value = `True`
	fmt.Println(yamlnode.IsBool(node))
	yamlnode.Print(node)
	// Output:
	// true
	// True
}

func ExampleIsBool_from_Yes() {
	node := new(yaml.Node)
	node.Kind = yaml.ScalarNode
	node.Tag = `!!bool`
	node.Value = `Yes`
	fmt.Println(yamlnode.IsBool(node))
	yamlnode.Print(node)
	// Output:
	// true
	// !!bool Yes
}
