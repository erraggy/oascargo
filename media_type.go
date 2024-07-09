package oascargo

// MediaType is a struct that provides a way to define media types for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#media-type-object
type MediaType struct {
	SpecificationExtensions
	// Schema is the schema defining the content of the request, response, or parameter.
	Schema Schema `json:"schema"`

	// Example is an example of the media type. The example object SHOULD be in the correct
	// format as specified by the media type. The Example field is mutually exclusive of the Examples field.
	// Furthermore, if referencing a schema which contains an example, the example value SHALL override
	// the example provided by the schema.
	Example any `json:"example"`

	// Examples of the media type. Each example object SHOULD match the media type and specified schema if present.
	// The Examples field is mutually exclusive of the Example field. Furthermore, if referencing a schema which
	// contains an example, the examples value SHALL override the example provided by the schema.
	Examples map[string]Example `json:"examples"`

	// Encoding is a map between a property name and its encoding information. The key, being the property name,
	// MUST exist in the schema as a property. The encoding object SHALL only apply to requestBody objects when
	// the media type is multipart or application/x-www-form-urlencoded.
	Encoding map[string]Encoding `json:"encoding"`
}
