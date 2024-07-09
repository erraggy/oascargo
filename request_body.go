package oascargo

// RequestBody is a struct that provides a way to define request bodies for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#request-body-object
type RequestBody struct {
	SpecificationExtensions
	// Ref is a string that provides a reference to a target object.
	// This MUST be in the form of a URI. When set this field is mutually exclusive
	// of ALL the other fields.
	Ref string `json:"$ref"`

	// Description is a brief description of the request body.
	Description string `json:"description,omitempty"`

	// Content describes the content of the request body. The key is a media type or media type range and the value
	// describes it. For requests that match multiple keys, only the most specific key is applicable. e.g. text/plain
	// overrides text/*
	// This field is REQUIRED.
	Content map[string]MediaType `json:"content"`

	// Required determines if the request body is required in the request. Defaults to false.
	Required bool `json:"required,omitempty"`
}
