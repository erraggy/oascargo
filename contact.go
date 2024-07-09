package oascargo

// Contact is a struct that provides contact information for the exposed API.
// see: https://spec.openapis.org/oas/v3.1.0.html#contact-object
type Contact struct {
	SpecificationExtensions
	// Name is a string that provides the identifying name of the contact person/organization.
	Name string `json:"name"`

	// URL is a string that provides the URL pointing to the contact information.
	// MUST be in the format of a URL.
	URL string `json:"url"`

	// Email is a string that provides the email address of the contact person/organization.
	// MUST be in the format of an email address.
	Email string `json:"email"`
}
