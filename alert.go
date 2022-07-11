package weatherkit

import "golang.org/x/text/language"

// WeatherAlertRequest requests weather alert details for a specific alert id.
type WeatherAlertRequest struct {
	// (Required) The unique identifier for the weather alert.
	ID string `json:"id,omitempty"`

	// (Required) The language tag to use for localizing responses.
	Language language.Tag `json:"language,omitempty"`
}

func (o WeatherAlertRequest) url() string {
	return weatherAlertEndpoint(o.Language, o.ID)
}

// WeatherAlertResponse contains an official message indicating severe weather from a reporting agency.
type WeatherAlertResponse struct {
	WeatherAlertData
	WeatherAlertSummary
}

// WeatherAlertData is the weather alert information.
type WeatherAlertData struct {
	// (Required) An object defining the geographic region the weather alert applies to.
	Area WeatherAlertArea `json:"area"`

	// (Required) An array of official text messages describing a severe weather event from the agency.
	EventText []EventText `json:"eventText"`
}

// EventText is the official text describing a severe weather event from the agency.
type EventText struct {
	// The ISO language code that the text is in.
	Language language.Tag `json:"language,omitempty"`

	// The severe weather event text.
	Text string `json:"text,omitempty"`
}

// WeatherAlertArea defines the geographic region the weather alert applies to.
type WeatherAlertArea struct {
	// Fields are undocumented.
}
