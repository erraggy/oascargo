package oascargo

// Schema is a struct that provides the definition of a schema.
// This object is a superset of the:
// JSON Schema Specification Draft 2020-12 (https://tools.ietf.org/html/draft-bhutton-json-schema-00)
//
// see also: https://spec.openapis.org/oas/v3.1.0.html#schema-object
type Schema struct {
	SpecificationExtensions
	// We first define the fields from the JSON Schema Specification Draft 2020-12.

	// Schema is a string that specifies the version of the JSON Schema Specification that is being used.
	// It is REQUIRED and MUST be the semantic version number of the JSON Schema Specification version that the JSON Schema
	// document uses.
	// versions: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.4
	Schema string `json:"$schema"`

	// Title is a string that provides the title of the JSON Schema.
	Title string `json:"title"`

	// Description is a string that provides a description of the JSON Schema.
	Description string `json:"description"`

	// Type is a string that provides the type of the JSON Schema.
	Type string `json:"type"`

	// Properties is a map of Schema objects that provides the properties of the JSON Schema.
	Properties map[string]Schema `json:"properties"`

	// Required is a list of strings that provides the required properties of the JSON Schema.
	Required []string `json:"required"`

	// Items is a Schema object that provides the items of the JSON Schema.
	Items *Schema `json:"items"`

	// AdditionalProperties is a Schema object that provides the additional properties of the JSON Schema.
	AdditionalProperties *Schema `json:"additionalProperties"`

	// Enum is a list of any that provides the enum of the JSON Schema.
	Enum []any `json:"enum"`

	// Format is a string that provides the format of the JSON Schema.
	Format string `json:"format"`

	// Ref is a string that provides the reference of the JSON Schema.
	Ref string `json:"$ref"`

	// Definitions is a map of Schema objects that provides the definitions of the JSON Schema.
	Definitions map[string]Schema `json:"definitions"`

	// AllOf is a list of Schema objects that provides the allOf of the JSON Schema.
	AllOf []Schema `json:"allOf"`

	// AnyOf is a list of Schema objects that provides the anyOf of the JSON Schema.
	AnyOf []Schema `json:"anyOf"`

	// We then define the fields from the OpenAPI Specification.

	// Example is an any that provides an example of the schema.
	// see: https://spec.openapis.org/oas/v3.1.0.html#schema-example
	Example any `json:"example"`

	// ExternalDocs is a struct that provides additional external documentation.
	// see: https://spec.openapis.org/oas/v3.1.0.html#external-documentation-object
	ExternalDocs ExternalDocumentation `json:"externalDocs"`

	// XML MAY be used only on properties schemas. It has no effect on root schemas. Adds additional metadata to
	// describe the XML representation of this property.
	// see: https://spec.openapis.org/oas/v3.1.0.html#xml-object
	XML XML `json:"xml"`

	// Discriminator adds support for polymorphism. The discriminator is an object name that is used to differentiate
	// between other schemas which may satisfy the payload description.
	// see: https://spec.openapis.org/oas/v3.1.0.html#discriminator-object
	Discriminator Discriminator `json:"discriminator"`
}

// Discriminator is a struct that provides support for polymorphism.
// The discriminator is an object name that is used to differentiate between other schemas which may satisfy the
// payload description.
// see: https://spec.openapis.org/oas/v3.1.0.html#discriminator-object
type Discriminator struct {
	SpecificationExtensions
	// PropertyName is a string that specifies the name of the property in the payload that will hold the
	// discriminator value.
	// The property name provided MUST be defined in the parent schema and it MUST be a required property.
	PropertyName string `json:"propertyName"`

	// Mapping is a map of string to string that provides a mapping between the discriminator value and schema name or
	// reference.
	// The discriminator value is used to select the schema definition in the payload using the value of the property
	// that is provided by propertyName.
	Mapping map[string]string `json:"mapping"`
}

// XML is a struct that provides additional information about the XML representation of the object.
// see: https://spec.openapis.org/oas/v3.1.0.html#xml-object
type XML struct {
	SpecificationExtensions
	// Name replaces the name of the element/attribute used for the described schema property.
	// When defined within items, it will affect the name of the individual XML elements within the list.
	// When defined alongside type being array (outside the items), it will affect the wrapping element and only if
	// wrapped is true. If wrapped is false, it will be ignored.
	Name string `json:"name"`

	// Namespace is a string that provides the URI of the namespace definition.
	// This MUST be in the form of an absolute URI.
	Namespace string `json:"namespace"`

	// Prefix is a string that provides the prefix for the name.
	Prefix string `json:"prefix"`

	// Attribute is a boolean that specifies whether the property definition translates to an attribute instead of an
	// element.
	Attribute bool `json:"attribute"`

	// Wrapped MAY be used only for an array definition. Signifies whether the array is wrapped (for example,
	// <books><book/><book/></books>) or unwrapped (<book/><book/>). Default value is false. The definition takes
	// effect only when defined alongside type being array (outside the items).
	Wrapped bool `json:"wrapped"`
}

// ExternalDocumentation is a struct that provides additional external documentation.
// see: https://spec.openapis.org/oas/v3.1.0.html#external-documentation-object
type ExternalDocumentation struct {
	SpecificationExtensions
	// Description is a string that provides a short description of the target documentation.
	Description string `json:"description"`

	// URL is a string that provides the URL for the target documentation.
	// It MUST be in the format of a URL.
	URL string `json:"url"`
}
