package yamlnode_test

import (
	"fmt"

	"github.com/rwxrob/yamlnode"
	"gopkg.in/yaml.v3"
)

func ExampleFromAsDoc() {
	node, err := yamlnode.FromAsDoc("foo")
	if err != nil {
		fmt.Println(err)
		return
	}

	// gets marshalled as same string as input
	yamlnode.Print(node)

	// content always encapsulated in a DocumentNode Kind
	fmt.Println(node.Kind == yaml.DocumentNode)

	// first in Content slice has value
	fmt.Println(node.Content[0].Kind == yaml.ScalarNode)

	// synonymous
	fmt.Println(node.Content[0].Tag)
	fmt.Println(node.Content[0].ShortTag())

	// long form as specified by YAML spec
	fmt.Println(node.Content[0].LongTag())

	// Output:
	// foo
	// true
	// true
	// !!str
	// !!str
	// tag:yaml.org,2002:str

}

func ExampleCast_bool_Tag_Still_String() {
	node, _ := yamlnode.From(`!!bool Yes`)
	fmt.Println(node.Tag)
	fmt.Println(node.Value)

	// Output:
	// !!str
	// !!bool Yes
}
