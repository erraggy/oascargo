package oascargo

// Tag adds metadata to a single tag that is used by the Operation Object.
// It is not mandatory to have a Tag Object per tag used there.
// see: https://spec.openapis.org/oas/v3.1.0.html#tag-object
type Tag struct {
	SpecificationExtensions
	// Name is the name of the tag.
	// It MUST be unique among all tags used in the OpenAPI document.
	// This field is REQUIRED.
	Name string `json:"name"`

	// Description for the tag.
	Description string `json:"description,omitempty"`

	// ExternalDocs is additional external documentation for this tag.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}
