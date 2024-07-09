package oascargo

// Example is a struct that provides a way to define examples for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#example-object
type Example struct {
	SpecificationExtensions
	// Summary is a short description for the example.
	Summary string `json:"summary"`

	// Description is a longer description for the example.
	Description string `json:"description"`

	// Ref is a string that provides a reference to an example object.
	// This MUST be in the form of a URI.
	// This field is mutually exclusive of the value and externalValue fields, and if set will override them.
	Ref string `json:"$ref"`

	// Value is the embedded literal example. The value field and externalValue field are mutually exclusive.
	// To represent examples of media types that cannot be naturally represented in JSON,
	// use a string value to contain the example, escaping where necessary.
	Value any `json:"value"`

	// ExternalValue is a URI that points to the literal example. This provides the capability to reference examples
	// that cannot easily be included in JSON documents. The value field and externalValue field are mutually exclusive.
	ExternalValue string `json:"externalValue"`
}
