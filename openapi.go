package oascargo

// OpenAPI is a struct that represents the root object of the OpenAPI document.
// see: https://spec.openapis.org/oas/v3.1.0.html#openapi-object
type OpenAPI struct {
	SpecificationExtensions
	// OpenAPI is a string that specifies the version of the OpenAPI Specification that is being used.
	// It is REQUIRED and MUST be the semantic version number of the OpenAPI Specification version that the OpenAPI
	// document uses.
	// versions: https://spec.openapis.org/oas/v3.1.0.html#versions
	OpenAPI string `json:"openapi"`

	// Info is a struct that provides metadata about the API.
	// see: https://spec.openapis.org/oas/v3.1.0.html#info-object
	Info Info `json:"info"`

	// JSONSchemaDialect is a string that specifies the default JSON Schema dialect to use for the API.
	// If omitted, "https://json-schema.org/draft/2020-12/schema" is assumed.
	JSONSchemaDialect string `json:"jsonSchemaDialect"`

	// Servers is a list of Server objects that provide connectivity information to a target server.
	// see: https://spec.openapis.org/oas/v3.1.0.html#server-object
	Servers []Server `json:"servers"`

	// Paths is a map of the available paths and operations for the API.
	Paths Paths `json:"paths"`

	// Webhooks is the incoming webhooks that MAY be received as part of this API and that the
	// API consumer MAY choose to implement. Closely related to the callbacks feature, this section
	// describes requests initiated other than by an API call, for example by an out of band registration.
	// The key name is a unique string to refer to each webhook, while the (optionally referenced) PathItem
	// Object describes a request that may be initiated by the API provider and the expected responses.
	// see: https://github.com/OAI/OpenAPI-Specification/tree/master/examples/v3.1/webhook-example.yaml
	Webhooks Paths `json:"webhooks,omitempty"`

	// Components is an element to hold various schemas for the document.
	Components *Components `json:"components,omitempty"`

	// Security is a declaration of which security mechanisms can be used across the API.
	// The list of values includes alternative security requirement objects that can be used.
	// Only one of the security requirement objects need to be satisfied to authorize a request.
	// To make security optional, an empty security requirement ({}) can be included in the array.
	Security []SecurityRequirement `json:"security,omitempty"`

	// Tags is a list of tags used by the document with additional metadata.
	// The order of the tags can be used to reflect on their order by the parsing tools.
	// Not all tags that are used by the Operation Object must be declared.
	// The tags that are not declared MAY be organized randomly or based on the tools' logic.
	// Each tag name in the list MUST be unique.
	Tags []Tag `json:"tags,omitempty"`

	// ExternalDocs is additional external documentation.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}
