package oascargo

// Info is a struct that provides metadata about the API.
// see: https://spec.openapis.org/oas/v3.1.0.html#info-object
type Info struct {
	SpecificationExtensions
	// Title is a string that provides the title of the API.
	// It is REQUIRED.
	Title string `json:"title"`

	// Summary is a string that provides a short summary of what the API does.
	Summary string `json:"summary"`

	// Description is a string that provides a longer description of the API.
	// It is REQUIRED.
	Description string `json:"description"`

	// TermsOfService is a string that provides the URL to the Terms of Service for the API.
	// This MUST be in the form of a URL.
	TermsOfService string `json:"termsOfService"`

	// Contact is a struct that provides contact information for the exposed API.
	// see: https://spec.openapis.org/oas/v3.1.0.html#contact-object
	Contact Contact `json:"contact"`

	// License is a struct that provides licensing information for the exposed API.
	// see: https://spec.openapis.org/oas/v3.1.0.html#license-object
	License License `json:"license"`

	// Version is a string that provides the version of the OpenAPI document.
	// It is REQUIRED.
	Version string `json:"version"`
}
