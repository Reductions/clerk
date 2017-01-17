package clerk

import "encoding/xml"

// A Channel defines an imput channle to be recorded
type Channel struct {
	Intermittent bool   // Defines whether the Channel is Intermittent
	Default      string // Default value of the Channle if Intermittent is true
	Name         string // Name of the channle
	Type         string // Type of the channle
	Parent       *Node  // Parent node.
	Mapping      string // Mapping type
}

// TraceNode is interface that gives the base for every element that could be grouped into a TraceGroups
type TraceNode interface {
}

// A TraceFormat is element that describes the format used to encode points
type TraceFormat struct {
	Children   []*Channel  // Child nodes.
	Attributes []*xml.Attr // Node attributes.
	Parent     *Node       // Parent node.
}

// A Trace is the basic element used to record the the trajectory of a pen
type Trace struct {
	Attributes []*xml.Attr // Node attributes.
	Parent     *Node       // Parent node.
	Value      string      // The value of the Trace
}

// A TraceGroup provides mechanisms to gather and combine traces into collections
type TraceGroup struct {
	Children   []*TraceNode // Child nodes.
	Attributes []*xml.Attr  // Node attributes.
	Parent     *Node        // Parent node.
}

// A TraceView element is used to include traces by reference
type TraceView struct {
	Children   []*TraceNode // Child nodes.
	Attributes []*xml.Attr  // Node attributes.
	Parent     *Node        // Parent node.
}
