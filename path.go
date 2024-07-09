package oascargo

// Paths is a map of URL paths to their PathItem objects.
type Paths struct {
	SpecificationExtensions
	Items map[string]PathItem `json:"-"`
}

// PathItem is a struct that provides a way to define paths and operations for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#path-item-object
type PathItem struct {
	SpecificationExtensions
	// Ref is a string that provides a reference to a target object.
	// This MUST be in the form of a URI. When set this field is mutually exclusive
	// of ALL the other fields.
	Ref string `json:"$ref"`

	// Summary is a string that provides a short summary of what the path item does.
	Summary string `json:"summary"`

	// Description is a string that provides a longer description of the path item.
	Description string `json:"description"`

	// GET is a definition of a GET operation on this path.
	GET *Operation `json:"get"`

	// PUT is a definition of a PUT operation on this path.
	PUT *Operation `json:"put"`

	// POST is a definition of a POST operation on this path.
	POST *Operation `json:"post"`

	// DELETE is a definition of a DELETE operation on this path.
	DELETE *Operation `json:"delete"`

	// OPTIONS is a definition of a OPTIONS operation on this path.
	OPTIONS *Operation `json:"options"`

	// HEAD is a definition of a HEAD operation on this path.
	HEAD *Operation `json:"head"`

	// PATCH is a definition of a PATCH operation on this path.
	PATCH *Operation `json:"patch"`

	// TRACE is a definition of a TRACE operation on this path.
	TRACE *Operation

	// Servers is an alternative server array to service all operations in this path.
	Servers []Server `json:"servers"`

	// Parameters is a list of parameters that are applicable for all the operations described under this path.
	// These parameters can be overridden at the operation level, but cannot be removed there.
	// The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination
	// of a name and location. The list can use the Reference Object to link to parameters that are defined at
	// the OpenAPI Object's components/parameters.
	Parameters []Parameter `json:"parameters"`
}

// CallBack is a map of possible out-of band callbacks related to the parent operation.
// Each value in the map is a PathItem Object that describes a set of requests that may be initiated by the API provider
// and the expected responses.
// see: https://spec.openapis.org/oas/v3.1.0.html#callback-object
type CallBack struct {
	SpecificationExtensions
	Items map[Expression]PathItem `json:"-"`
}
