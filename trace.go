package clerk

import "encoding/xml"

// A Channle defines an imput channle to be recorded
type Channel struct {
	Intermittent bool   // Defines whether the Channel
	Default      string // Default value of the Channle if Intermittent is true
	Name         string // Name of the channle
	Type         string // Type of the channle
	Parent       *Node  // Parent node.
	Mapping      string // Mapping type
}

// A TraceFormat is element that describes the format used to encode points
type TraceFormat struct {
	Children   []*Channel  // Child nodes.
	Attributes []*xml.Attr // Node attributes.
	Parent     *Node       // Parent node.
}
