package oascargo

// SpecificationExtensions is a struct that provides a way to add additional metadata to the OpenAPI Specification.
// It is an object where the keys are the name of the extension and the values are any JSON-compatible value.
// The value of the extensions field MUST be an object.
// Field names beginning with "x-" are reserved for extensions.
// see: https://spec.openapis.org/oas/v3.1.0.html#specification-extensions
type SpecificationExtensions map[string]any
