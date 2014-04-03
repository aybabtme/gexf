package gexf

// GEXFType is an type for values Nodes can hold.
type GEXFType string

// All the recognized GEXTTypes.
const (
	Long       GEXFType = "long"
	Double     GEXFType = "double"
	FLoat      GEXFType = "float"
	Boolean    GEXFType = "boolean"
	ListString GEXFType = "liststring"
	String     GEXFType = "string"
	AnyURI     GEXFType = "anyURI"
)
