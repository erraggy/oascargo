package oascargo

// Encoding is a struct that provides a way to define encoding for an OpenAPI Specification.
// see: https://spec.openapis.org/oas/v3.1.0.html#encoding-object
type Encoding struct {
	SpecificationExtensions
	// ContentType for encoding a specific property. Default value depends on the property type:
	// for object - application/json; for array – the default is defined based on the inner type;
	// for all other cases the default is application/octet-stream.
	// The value can be a specific media type (e.g. application/json),
	// a wildcard media type (e.g. image/*), or a comma-separated list of the two types.
	ContentType string `json:"contentType"`

	// Headers is a map allowing additional information to be provided as headers,
	// for example Content-Disposition. Content-Type is described separately and SHALL be ignored
	// in this section. This property SHALL be ignored if the request body media type is not a multipart.
	Headers map[string]Header `json:"headers"`

	// Style describes how the value is serialized depending on the type of the parameter value.
	// See Parameter Object (https://spec.openapis.org/oas/v3.1.0.html#parameter-object)
	// for details on the style property. The behavior follows the same values as query parameters,
	// including default values. This property SHALL be ignored if the request body media type is not
	// application/x-www-form-urlencoded or multipart/form-data.
	// If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	Style string `json:"style"`

	// Explode when this is true, property values of type array or object generate separate parameters
	// for each value of the array, or key-value pair of the map. For other types of properties this property
	// has no effect. When style is form, the default value is true. For all other styles, the default value is false.
	// This property SHALL be ignored if the request body media type is not
	// application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of
	// contentType (implicit or explicit) SHALL be ignored.
	Explode bool `json:"explode"`

	// AllowReserved determines whether the parameter value SHOULD allow reserved characters,
	// as defined by RFC3986 :/?#[]@!$&'()*+,;= to be included without percent-encoding.
	// The default value is false. This property SHALL be ignored if the request body media type is not
	// application/x-www-form-urlencoded or multipart/form-data. If a value is explicitly defined, then the value of
	// contentType (implicit or explicit) SHALL be ignored.
	AllowReserved bool `json:"allowReserved"`
}
