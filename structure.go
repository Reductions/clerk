/*
Package clerk is a go library for working with InkML dataformat
*/
package clerk

import (
	"encoding/xml"
	"image"
)

// A Document is the wrapping type of any InkML
type Document struct {
	Version string
	Root    *Ink
}

// Node is interface that gives the base for all valid InkML tags
type Node interface {
	Parent() *Node
	Children() []*Node
	AddChild(newChild *Node) error
	AddChildAfter(newChild *Node, refChild *Node) error
	AddChildBefor(newChild *Node, refChild *Node) error
}

// An Ink element is the root element of any InkML instance
type Ink struct {
	Children   []*Node     // Child nodes.
	Attributes []*xml.Attr // Node attributes.
}

// New creates InkML document with just Ink tag inside it
func New() *Document {
	return nil
}

// ParseBytes creates an InkML document from byte slice
func ParseBytes(bytes []byte) (*Document, error) {
	return nil, nil
}

// ToBytes serialises the dicument as byte slice
func (*Document) ToBytes() []byte {
	return nil
}

// ToImage writes the InkML to and Image
func (*Document) ToImage(img *image.Image) error {
	return nil
}
