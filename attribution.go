package weatherkit

// AttributionRequest requests attribution details.
type AttributionRequest struct {
	// (Required) The language tag to use for localizing responses.
	Language string `json:"language,omitempty"`
}

func (o AttributionRequest) url() string {
	return attribution(o.Language)
}

// AttributionResponse contains an official attribution branding.
type AttributionResponse struct {
	LogoDark1x   string `json:"logoDark@1x,omitempty"`
	LogoDark2x   string `json:"logoDark@2x,omitempty"`
	LogoDark3x   string `json:"logoDark@3x,omitempty"`
	LogoLight1x  string `json:"logoLight@1x,omitempty"`
	LogoLight2x  string `json:"logoLight@2x,omitempty"`
	LogoLight3x  string `json:"logoLight@3x,omitempty"`
	LogoSquare1x string `json:"logoSquare@1x,omitempty"`
	LogoSquare2x string `json:"logoSquare@2x,omitempty"`
	LogoSquare3x string `json:"logoSquare@3x,omitempty"`
	ServiceName  string `json:"serviceName,omitempty"`
}
