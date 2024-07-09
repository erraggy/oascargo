package oascargo

// Parameter is a struct that provides a way to define parameters for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#parameter-object
type Parameter struct {
	SpecificationExtensions
	// Ref is a string that provides a reference to a target object.
	// This MUST be in the form of a URI. When set this field is mutually exclusive
	// of ALL the other fields.
	Ref string `json:"$ref"`

	// Name is a string that provides the name of the parameter.
	// Parameter names are case sensitive.
	// This field is REQUIRED.
	Name string `json:"name"`

	// In is a string that provides the location of the parameter.
	// Possible values are "query", "header", "path", "cookie".
	// This field is REQUIRED.
	In string `json:"in"`

	// Description is a string that provides a brief description of the parameter.
	Description string `json:"description"`

	// Required is a boolean that determines whether this parameter is mandatory.
	// If the parameter location is "path", this property is REQUIRED and its value MUST be true.
	// Otherwise, the property MAY be included and its default value is false.
	Required bool `json:"required"`

	// Deprecated is a boolean that specifies if a parameter is deprecated
	// and SHOULD be transitioned out of usage. Default value is false.
	Deprecated bool `json:"deprecated"`

	// AllowEmptyValue sets the ability to pass empty-valued parameters.
	// This is valid only for query parameters and allows sending a parameter with an empty value.
	// Default value is false.
	// If style is used, and if behavior is n/a (cannot be serialized),
	// the value of allowEmptyValue SHALL be ignored. Use of this property is NOT RECOMMENDED,
	// as it is likely to be removed in a later revision.
	AllowEmptyValue bool `json:"allowEmptyValue"`

	// Style describes how the parameter value will be serialized depending on the type of the parameter value.
	// Default values (based on value of in): for query - form; for path - simple;
	// for header - simple; for cookie - form.
	Style string `json:"style"`

	// Explode when this is true, parameter values of type array or object generate separate
	// parameters for each value of the array or key-value pair of the map.
	// For other types of parameters this property has no effect.
	// When style is form, the default value is true.
	// For all other styles, the default value is false.
	Explode bool `json:"explode"`

	// AllowReserved determines whether the parameter value SHOULD allow reserved characters,
	// as defined by RFC3986 :/?#[]@!$&'()*+,;= to be included without percent-encoding.
	// This property only applies to parameters with an in value of query.
	// The default value is false.
	AllowReserved bool `json:"allowReserved"`

	// Schema the schema defining the type used for the parameter.
	Schema *Schema `json:"schema"`

	// Example of the parameter's potential value. The example SHOULD match the specified schema and encoding properties
	// if present. The Example field is mutually exclusive of the Examples field. Furthermore, if referencing a schema
	// which contains an example, the example value SHALL override the example provided by the schema. To represent
	// examples of media types that cannot naturally be represented in JSON or YAML, a string value can contain the
	// example with escaping where necessary.
	Example any `json:"example"`

	// Examples is a map containing the examples of the parameter. The Examples field is mutually exclusive of the
	// Example field. Furthermore, if referencing a schema which contains an example, the examples value SHALL override
	// the example provided by the schema.
	Examples map[string]Example `json:"examples"`

	// Content is a map containing the representations for the parameter. The key is the media type and the value
	// describes it. The map MUST only contain one entry.
	Content map[string]MediaType `json:"content"`
}
