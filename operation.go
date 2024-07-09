package oascargo

// Operation is a struct that provides a way to define paths and operations for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#operation-object
type Operation struct {
	SpecificationExtensions
	// Tags is a list of tags for API documentation control. Tags can be used for logical grouping of
	// operations by resources or any other qualifier.
	Tags []string `json:"tags,omitempty"`

	// Summary is a string that provides a short summary of what the operation does.
	Summary string `json:"summary,omitempty"`

	// Description is a string that provides a longer description of the operation.
	Description string `json:"description,omitempty"`

	// ExternalDocs may provide additional external documentation for this operation.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`

	// OperationID is a unique string used to identify the operation. The id MUST be unique among all operations
	// described in the API. The operationId value is case-sensitive. Tools and libraries MAY use the operationId to
	// uniquely identify an operation, therefore, it is RECOMMENDED to follow common programming naming conventions.
	OperationID string `json:"operationId"`

	// Parameters is a list of parameters that are applicable for this operation. If a parameter is already defined at
	// the Path Item, the new definition will override it but can never remove it.
	Parameters []Parameter `json:"parameters,omitempty"`

	// RequestBody is a literal value or runtime Expression that specifies the value of the request body.
	RequestBody *RequestBody `json:"requestBody,omitempty"`

	// Responses is a list of possible responses as they are returned from executing this operation.
	Responses *Responses `json:"responses,omitempty"`

	// Callbacks is a map of possible out-of band callbacks related to the parent operation.
	// The key is a unique identifier for the Callback Object. Each value in the map is a Path Item Object that
	// describes a set of requests that may be initiated by the API provider and the expected responses.
	Callbacks map[string]CallBack `json:"callbacks,omitempty"`

	// Deprecated declares this operation to be deprecated.
	// Consumers SHOULD refrain from usage of the declared operation.
	// Default value is false.
	Deprecated bool `json:"deprecated,omitempty"`

	// Security is a declaration of which security mechanisms can be used for this operation.
	// The list of values includes alternative security requirement objects that can be used.
	// Only one of the security requirement objects need to be satisfied to authorize a request.
	// To make security optional, an empty security requirement ({}) can be included in the array.
	// This definition overrides any declared top-level security.
	// To remove a top-level security declaration, an empty array can be used.
	Security []SecurityRequirement `json:"security,omitempty"`

	// Servers is an alternative server array to service this operation. If an alternative server object is specified
	// at the PathItem Object or Root level, it will be overridden by this value.
	Servers []Server `json:"servers,omitempty"`
}
