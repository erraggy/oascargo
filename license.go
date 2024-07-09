package oascargo

// License is a struct that provides licensing information for the exposed API.
// see: https://spec.openapis.org/oas/v3.1.0.html#license-object
type License struct {
	SpecificationExtensions
	// Name is a string that provides the license name used for the API.
	// It is REQUIRED.
	Name string `json:"name"`

	// Identifier is a string that provides an SPDX license expression for the API.
	// The identifier field is mutually exclusive of the url field.
	Identifier string `json:"identifier"`

	// URL is a string that provides the URL pointing to the license used for the API.
	// MUST be in the format of a URL.
	URL string `json:"url"`
}
