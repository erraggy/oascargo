package oascargo

// Components holds a set of reusable objects for different aspects of the OAS.
// All objects defined within the components object will have no effect on the API
// unless they are explicitly referenced from properties outside the components object.
// see: https://spec.openapis.org/oas/v3.1.0.html#components-object
type Components struct {
	SpecificationExtensions
	// Schemas is an object to hold reusable Schema Objects.
	Schemas map[string]Schema `json:"schemas,omitempty"`

	// Responses is an object to hold reusable Response Objects.
	Responses map[string]Response `json:"responses,omitempty"`

	// Parameters is an object to hold reusable Parameter Objects.
	Parameters map[string]Parameter `json:"parameters,omitempty"`

	// Examples is an object to hold reusable Example Objects.
	Examples map[string]Example `json:"examples,omitempty"`

	// RequestBodies is an object to hold reusable Request Body Objects.
	RequestBodies map[string]RequestBody `json:"requestBodies,omitempty"`

	// Headers is an object to hold reusable Header Objects.
	Headers map[string]Header `json:"headers,omitempty"`

	// SecuritySchemes is an object to hold reusable Security Scheme Objects.
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`

	// Links is an object to hold reusable Link Objects.
	Links map[string]Link `json:"links,omitempty"`

	// Callbacks is an object to hold reusable Callback Objects.
	Callbacks map[string]CallBack `json:"callbacks,omitempty"`

	// PathItems is an object to hold reusable Path Item Objects.
	PathItems map[string]PathItem `json:"pathItems,omitempty"`
}
