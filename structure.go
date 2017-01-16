/*
Package clerk is a go library for working with InkML dataformat
*/
package clerk

import "encoding/xml"

// Node is interface that gives the base for all valid InkML tags
type Node interface {
}

// An Ink element is the root element of any InkML instance
type Ink struct {
	Children   []*Node     // Child nodes.
	Attributes []*xml.Attr // Node attributes.
}
