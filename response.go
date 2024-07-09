package oascargo

// Response is a struct that provides a way to define responses for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#response-object
type Response struct {
	SpecificationExtensions
	// Ref is a string that provides a reference to a target object.
	// This MUST be in the form of a URI. When set this field is mutually exclusive
	// of ALL the other fields.
	Ref string `json:"$ref"`

	// Description is a string that provides a description of the response.
	Description string `json:"description"`

	// Headers describes the headers that can be sent as part of the response.
	Headers map[string]Header `json:"headers,omitempty"`

	// Content is a map containing descriptions of potential response payloads.
	// The key is a media type or media type range and the value describes it.
	// For responses that match multiple keys, only the most specific key is applicable.
	// e.g. text/plain overrides text/*
	Content map[string]MediaType `json:"content"`

	// Links is a map of links that can be followed from the response.
	// The key of the map is a short name for the link following the naming constraints of the names for
	// Component Objects (https://spec.openapis.org/oas/v3.1.0.html#components-object)
	Links map[string]Link `json:"links,omitempty"`
}

// Responses is a container for the expected responses of an operation. The container maps a HTTP response
// status code to the expected response.
type Responses struct {
	SpecificationExtensions
	// Default is the documentation of responses other than the ones declared for specific HTTP response codes.
	// Use this field to cover undeclared responses.
	Default *Response `json:"default,omitempty"`

	// Responses is a map of HTTP response status codes, where the key is the status code and the value is the response.
	// Any HTTP status code can be used as the property name, but only one property per code,
	// to describe the expected response for that HTTP status code. This field MUST be enclosed in quotation marks
	// (for example, “200”) for compatibility between JSON and YAML. To define a range of response codes,
	// this field MAY contain the uppercase wildcard character X. For example, 2XX represents all response codes
	// between [200-299]. Only the following range definitions are allowed: 1XX, 2XX, 3XX, 4XX, and 5XX.
	// If a response is defined using an explicit code, the explicit code definition takes precedence over
	// the range definition for that code.
	Responses map[string]Response `json:"responses,omitempty"`
}
