package yamlnode

import "gopkg.in/yaml.v3"

// MapKeyFor returns from.Content[contentIndex-1].Value and should only be
// passed a valnode that passes [IsMapVal]. Returns empty string if
// unable to lookup.
func MapKeyFor(valnode, from *yaml.Node, contentIndex int) string {
	if valnode == nil || from == nil || from.Content == nil {
		return ""
	}
	return from.Content[contentIndex-1].Value
}

// MapValFor returns from.Content[contentIndex+1].Value and should only be
// called passed a keynode that passes [IsMapKey]. Returns empty string if
// unable to lookup.
func MapValFor(keynode, from *yaml.Node, contentIndex int) string {
	if keynode == nil || from == nil || from.Content == nil {
		return ""
	}
	return from.Content[contentIndex-1].Value
}
