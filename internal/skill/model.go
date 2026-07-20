package skill

type Resource struct {
	Name string `json:"name"`
	Type string `json:"type"` // worskspace, filepath, url

	Path string `json:"path,omitempty"` // filepath , folder path
	URL  string `json:"url,omitempty"`  // url
}

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	Permissions []string `json:"permissions"`
	Guidelines  []string `json:"guidelines"`
	Constraints []string `json:"constraints"`

	Tools     []string `json:"tools"`
	Resources []string `json:"resources"`
}

type RetryPolicy struct {
	MaxAttempts int `json:"maxAttempts"`
}

type ErrorHandler struct {
	Strategy string `json:"strategy"`
}

type Operation struct {
	Metadata Metadata `json:"metadata"`

	Dependencies []string `json:"dependencies"`

	Inputs  []string `json:"inputs"`
	Outputs []string `json:"outputs"`

	Examples   []string `json:"examples"`
	Validation []string `json:"validation"`

	RetryPolicy  RetryPolicy  `json:"retryPolicy"`
	ErrorHandler ErrorHandler `json:"errorHandler"`
}

type Chain struct {
	Metadata Metadata `json:"metadata"`

	Operations []Operation `json:"operations"`
}

type Hooks struct {
	Before []string `json:"before"`
	After  []string `json:"after"`
}

type Skill struct {
	Metadata Metadata `json:"metadata"`

	Chains []Chain `json:"chains"`
	Hooks  Hooks   `json:"hooks"`
}
