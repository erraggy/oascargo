package oascargo

// Link is a struct that provides a way to define links for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#link-object
type Link struct {
	SpecificationExtensions
	// Ref is a string that provides a reference to a target object.
	// This MUST be in the form of a URI. When set this field is mutually exclusive
	// of ALL the other fields.
	Ref string `json:"$ref"`

	// OperationRef is a string that provides a reference to an operation in the spec.
	// It may be a relative or an absolute URI reference to an OAS operation.
	// This field is mutually exclusive of the OperationID field.
	OperationRef string `json:"operationRef"`

	// OperationID is the name of an existing, resolvable OAS operation, as defined with a unique operationId.
	// This field is mutually exclusive of the OperationRef field.
	OperationID string `json:"operationId"`

	// Parameters is a map of runtime expressions for selective substitution.
	// The key is the name of a parameter from the operation's parameters.
	// The value can be a constant or an expression to be evaluated and passed to the linked operation.
	Parameters map[string]Expression `json:"parameters,omitempty"`

	// RequestBody is a literal value or runtime Expression that specifies the value of the request body.
	RequestBody Expression `json:"requestBody,omitempty"`

	// Description is a string that provides a description of the link.
	Description string `json:"description,omitempty"`

	// Server is a server object to be used by the target operation.
	Server *Server `json:"server,omitempty"`
}

// Expression is a string that provides a runtime expression to be evaluated in the context of the OAS operation.
// see: https://spec.openapis.org/oas/v3.1.0.html#runtime-expressions
type Expression string
