package oascargo

// Header is a struct that provides a way to define headers for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#header-object
type Header struct {
	SpecificationExtensions
	// Ref is a string that provides a reference to a Header object.
	// This MUST be in the form of a URI.
	// This field is mutually exclusive of ALL OTHER FIELDS, and if set will override them.
	Ref string `json:"$ref"`

	// Description is a string that provides a brief description of the header.
	Description string `json:"description"`

	// Required is a boolean that determines whether this header is mandatory.
	// The property MAY be included and its default value is false.
	Required bool `json:"required"`

	// Deprecated is a boolean that specifies if a header is deprecated
	// and SHOULD be transitioned out of usage. Default value is false.
	Deprecated bool `json:"deprecated"`

	// Style describes how the header value will be serialized depending on the type of the parameter value.
	// Default value: simple
	Style string `json:"style"`

	// Explode when this is true, header values of type array or object generate separate
	// headers for each value of the array or key-value pair of the map.
	// For other types of parameters this property has no effect.
	// When style is form, the default value is true.
	// For all other styles, the default value is false.
	Explode bool `json:"explode"`

	// Schema the schema defining the type used for the parameter.
	Schema *Schema `json:"schema"`

	// Example of the header's potential value. The example SHOULD match the specified schema and encoding properties
	// if present. The Example field is mutually exclusive of the Examples field. Furthermore, if referencing a schema
	// which contains an example, the example value SHALL override the example provided by the schema. To represent
	// examples of media types that cannot naturally be represented in JSON or YAML, a string value can contain the
	// example with escaping where necessary.
	Example any `json:"example"`

	// Examples is a map containing the examples of the header. The Examples field is mutually exclusive of the
	// Example field. Furthermore, if referencing a schema which contains an example, the examples value SHALL override
	// the example provided by the schema.
	Examples map[string]Example `json:"examples"`

	// Content is a map containing the representations for the header. The key is the media type and the value
	// describes it. The map MUST only contain one entry.
	Content map[string]MediaType `json:"content"`
}
