package oascargo

// Server is a struct that provides connectivity information to a target server.
// see: https://spec.openapis.org/oas/v3.1.0.html#server-object
type Server struct {
	SpecificationExtensions
	// URL is a string that provides the URL to the target server.
	// It is REQUIRED.
	URL string `json:"url"`

	// Description is a string that provides a description for the server.
	Description string `json:"description"`

	// Variables is a map of ServerVariable objects that can be used for substitution in the server URL template.
	Variables map[string]ServerVariable `json:"variables"`
}

// ServerVariable is a struct that provides connectivity information to a target server.
// see: https://spec.openapis.org/oas/v3.1.0.html#server-variable-object
type ServerVariable struct {
	SpecificationExtensions
	// Enum is a list of string values that this server uses.
	// If the enum is defined, and the variable is in use, the value MUST exist in the enum.
	Enum []string `json:"enum"`

	// Default is a string that specifies the default value to use for substitution, and to send,
	// If an alternate value is not supplied.
	// If the enum is defined, the value SHOULD exist in the enum.
	Default string `json:"default"`

	// Description is a string that provides a description for the server variable.
	Description string `json:"description"`
}
