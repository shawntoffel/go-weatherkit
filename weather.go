package weatherkit

import (
	"net/url"
	"time"

	"golang.org/x/text/language"
)

// The date time format to use in query parameters.
const dateTimeFormat = time.RFC3339

// WeatherRequest obtains weather data for the specified location.
type WeatherRequest struct {
	// The language tag to use for localizing responses.
	Language language.Tag

	// The latitude of the desired location.
	Latitude float64

	// The longitude of the desired location.
	Longitude float64

	// The ISO Alpha-2 country code for the requested location.
	// This parameter is necessary for air quality and weather alerts.
	CountryCode string

	// The time to obtain current conditions. Defaults to now.
	CurrentAsOf *time.Time

	// The time to end the daily forecast.
	// If this parameter is absent, daily forecasts run for 10 days.
	DailyEnd *time.Time

	// The time to start the daily forecast.
	// If this parameter is absent, daily forecasts start on the current day.
	DailyStart *time.Time

	// A list of data sets to include in the response.
	DataSets DataSets

	// The time to end the hourly forecast.
	// If this parameter is absent, hourly forecasts run 24 hours or the length of the daily forecast, whichever is longer.
	HourlyEnd *time.Time

	// The time to start the hourly forecast.
	// If this parameter is absent, hourly forecasts start on the current hour.
	HourlyStart *time.Time

	// (Required) The name of the timezone to use for rolling up weather forecasts into daily forecasts.
	Timezone string
}

func (o WeatherRequest) url() string {
	q := url.Values{}

	if o.CountryCode != "" {
		q.Add("countryCode", o.CountryCode)
	}
	if o.CurrentAsOf != nil {
		q.Add("currentAsOf", o.CurrentAsOf.Format(dateTimeFormat))
	}
	if o.DailyEnd != nil {
		q.Add("dailyEnd", o.DailyEnd.Format(dateTimeFormat))
	}
	if o.DailyStart != nil {
		q.Add("dailyStart", o.DailyStart.Format(dateTimeFormat))
	}
	if len(o.DataSets) > 0 {
		q.Add("dataSets", o.DataSets.String())
	}
	if o.HourlyEnd != nil {
		q.Add("hourlyEnd", o.HourlyEnd.Format(dateTimeFormat))
	}
	if o.HourlyStart != nil {
		q.Add("hourlyStart", o.HourlyStart.Format(dateTimeFormat))
	}
	if o.Timezone != "" {
		q.Add("timezone", o.Timezone)
	}

	return weatherEndpoint(o.Language, o.Latitude, o.Longitude, q)
}

// WeatherResponse contains all requested properties.
type WeatherResponse struct {
	// The current weather for the requested location.
	CurrentWeather *CurrentWeather `json:"currentWeather,omitempty"`

	// The daily forecast for the requested location.
	ForcastDaily *DailyForecast `json:"forecastDaily,omitempty"`

	// The hourly forecast for the requested location.
	ForcastHourly *HourlyForecast `json:"forecastHourly,omitempty"`

	// The next hour forecast for the requested location.
	ForcastNextHour *NextHourForecast `json:"forecastNextHour,omitempty"`

	// Weather alerts for the requested location.
	WeatherAlerts *WeatherAlertCollection `json:"weatherAlerts,omitempty"`
}

// PrecipitationType is the type of precipitation forecasted to occur during the day.
type PrecipitationType string

const (
	// No precipitation is occurring.
	PrecipitationTypeClear PrecipitationType = "clear"

	// An unknown type of precipitation is occuring.
	PrecipitationTypePrecipitation PrecipitationType = "precipitation"

	// Rain or freezing rain is falling.
	PrecipitationTypeRain PrecipitationType = "rain"

	// Snow is falling.
	PrecipitationTypeSnow PrecipitationType = "snow"

	// Sleet or ice pellets are falling.
	PrecipitationTypeSleet PrecipitationType = "sleet"

	// Hail is falling.
	PrecipitationTypeHail PrecipitationType = "hail"

	// Winter weather (wintery mix or wintery showers) is falling.
	PrecipitationTypeMixed PrecipitationType = "mixed"
)

// PressureTrend is the direction of change of the sea level air pressure.
type PressureTrend string

const (
	// The sea level air pressure is increasing.
	PressureTrendRising PressureTrend = "rising"

	// The sea level air pressure is decreasing.
	PressureTrendFalling PressureTrend = "falling"

	// The sea level air pressure is remaining about the same.
	PressureTrendSteady PressureTrend = "steady"
)

// HourlyForecast represents the various weather phenomena occurring over a period of time
type HourlyForecast struct {
	ProductData

	// The hourly forecast information.
	Hours []HourWeatherConditions `json:"hours,omitempty"`
}

// HourWeatherConditions contains the historical or forecasted weather conditions for a specified hour.
type HourWeatherConditions struct {
	// (Required) The percentage of the sky covered with clouds during the period, from 0 to 1.
	CloudCover float64 `json:"cloudCover"`

	// (Required) An enumeration value indicating the condition at the time.
	ConditionCode string `json:"conditionCode,omitempty"`

	// Indicates whether the hour starts during the day or night.
	DayLight bool `json:"daylight"`

	// (Required) The starting date and time of the forecast.
	ForecastStart *time.Time `json:"forecastStart,omitempty"`

	// (Required) The relative humidity at the start of the hour, from 0 to 1.
	Humidity float64 `json:"humidity"`

	// (Required) The chance of precipitation forecasted to occur during the hour, from 0 to 1.
	PrecipitationChance float64 `json:"precipitationChance"`

	// (Required) The type of precipitation forecasted to occur during the period.
	PrecipitationType PrecipitationType `json:"precipitationType,omitempty"`

	// (Required) The sea-level air pressure, in millibars.
	Pressure float64 `json:"pressure"`

	// The direction of change of the sea-level air pressure.
	PressureTrend PressureTrend `json:"pressureTrend"`

	// The rate at which snow crystals are falling, in millimeters per hour.
	SnowfallIntensity float64 `json:"snowfallIntensity"`

	// (Required) The temperature at the start of the hour, in degrees Celsius.
	Temperature float64 `json:"temperature"`

	// (Required) The feels-like temperature when considering wind and humidity, at the start of the hour, in degrees Celsius.
	TemperatureApparent float64 `json:"temperatureApparent"`

	// The temperature at which relative humidity is 100% at the top of the hour, in degrees Celsius.
	TemperatureDewPoint float64 `json:"temperatureDewPoint"`

	// (Required) The level of ultraviolet radiation at the start of the hour.
	UvIndex int64 `json:"uvIndex"`

	// (Required) The distance at which terrain is visible at the start of the hour, in meters.
	Visibility float64 `json:"visibility"`

	// The direction of the wind at the start of the hour, in degrees.
	WindDirection float64 `json:"windDirection"`

	// The maximum wind gust speed during the hour, in kilometers per hour.
	WindGust float64 `json:"windGust"`

	// (Required) The wind speed at the start of the hour, in kilometers per hour.
	WindSpeed float64 `json:"windSpeed"`

	// The amount of precipitation forecasted to occur during period, in millimeters.
	PrecipitationAmount float64 `json:"precipitationAmount"`
}

// MoonPhase is the shape of the moon as seen by an observer on the ground at a given time.
type MoonPhase string

const (
	// The moon isn’t visible.
	MoonPhaseNew = "new"

	// A crescent-shaped sliver of the moon is visible, and increasing in size.
	MoonPhaseWaxingCrescent = "waxingCrescent"

	// Approximately half of the moon is visible, and increasing in size.
	MoonPhaseFirstQuarter = "firstQuarter"

	// The entire disc of the moon is visible.
	MoonPhaseFull = "full"

	// More than half of the moon is visible, and increasing in size.
	MoonPhaseWaxingGibbous = "waxingGibbous"

	// More than half of the moon is visible, and decreasing in size.
	MoonPhaseWaningGibbous = "waningGibbous"

	// Approximately half of the moon is visible, and decreasing in size.
	MoonPhaseThirdQuarter = "thirdQuarter"

	// A crescent-shaped sliver of the moon is visible, and decreasing in size.
	MoonPhaseWaningCrescent = "waningCrescent"
)

// HourlyForecast represents the various weather phenomena occurring over a period of time
type DailyForecast struct {
	ProductData
	Days []DayWeatherConditions `json:"days,omitempty"`
}

// DayWeatherConditions contains the historical or forecasted weather conditions for a specified day.
type DayWeatherConditions struct {
	// (Required) An enumeration value indicating the condition at the time.
	ConditionCode string `json:"conditionCode,omitempty"`

	// The forecast between 7 AM and 7 PM for the day.
	DaytimeForecast DayPartForecast `json:"daytimeForecast,omitempty"`

	// (Required) The ending date and time of the day.
	ForecastEnd *time.Time `json:"forecastEnd,omitempty"`

	// (Required) The starting date and time of the day.
	ForecastStart *time.Time `json:"forecastStart,omitempty"`

	// (Required) The maximum ultraviolet index value during the day.
	MaxUvIndex int64 `json:"maxUvIndex"`

	// (Required) The phase of the moon on the specified day.
	MoonPhase MoonPhase `json:"moonPhase,omitempty"`

	// The time of moonrise on the specified day.
	MoonRise *time.Time `json:"moonrise,omitempty"`

	// The time of moonset on the specified day.
	MoonSet *time.Time `json:"moonset,omitempty"`

	// The day part forecast between 7 PM and 7 AM for the overnight.
	OvernightForecast DayPartForecast `json:"overnightForecast,omitempty"`

	// (Required) The amount of precipitation forecasted to occur during the day, in millimeters.
	PrecipitationAmount float64 `json:"precipitationAmount"`

	// (Required) The chance of precipitation forecasted to occur during the day.
	PrecipitationChance float64 `json:"precipitationChance"`

	// (Required) The type of precipitation forecasted to occur during the day.
	PrecipitationType PrecipitationType `json:"precipitationType,omitempty"`

	// (Required) The depth of snow as ice crystals forecasted to occur during the day, in millimeters.
	SnowfallAmount float64 `json:"snowfallAmount"`

	// The time when the sun is lowest in the sky.
	SolarMidnight *time.Time `json:"solarMidnight,omitempty"`

	// The time when the sun is highest in the sky.
	SolarNoon *time.Time `json:"solarNoon,omitempty"`

	// The time when the top edge of the sun reaches the horizon in the morning.
	Sunrise *time.Time `json:"sunrise,omitempty"`

	// The time when the sun is 18 degrees below the horizon in the morning.
	SunriseAstronomical *time.Time `json:"sunriseAstronomical,omitempty"`

	// The time when the sun is 6 degrees below the horizon in the morning.
	SunriseCivil *time.Time `json:"sunriseCivil,omitempty"`

	// The time when the sun is 12 degrees below the horizon in the morning.
	SunriseNautical *time.Time `json:"sunriseNautical,omitempty"`

	// The time when the top edge of the sun reaches the horizon in the evening.
	Sunset *time.Time `json:"sunset,omitempty"`

	// The time when the sun is 18 degrees below the horizon in the evening.
	SunsetAstronomical *time.Time `json:"sunsetAstronomical,omitempty"`

	// The time when the sun is 6 degrees below the horizon in the evening.
	SunsetCivil *time.Time `json:"sunsetCivil,omitempty"`

	// The time when the sun is 12 degrees below the horizon in the evening.
	SunsetNautical *time.Time `json:"sunsetNautical,omitempty"`

	// (Required) The maximum temperature forecasted to occur during the day, in degrees Celsius.
	TemperatureMax float64 `json:"temperatureMax"`

	// (Required) The minimum temperature forecasted to occur during the day, in degrees Celsius.
	TemperatureMin float64 `json:"temperatureMin"`
}

// DayPartForecast is a summary forecast for a daytime or overnight period.
type DayPartForecast struct {
	// (Required) The percentage of the sky covered with clouds during the period, from 0 to 1.
	CloudCover float64 `json:"cloudCover"`

	// (Required) An enumeration value indicating the condition at the time.
	ConditionCode string `json:"conditionCode,omitempty"`

	// (Required) The ending date and time of the forecast.
	ForecastEnd *time.Time `json:"forecastEnd,omitempty"`

	// (Required) The starting date and time of the forecast.
	ForecastStart *time.Time `json:"forecastStart,omitempty"`

	// (Required) The relative humidity during the period, from 0 to 1.
	Humidity float64 `json:"humidity"`

	// (Required) The amount of precipitation forecasted to occur during the period, in millimeters.
	PrecipitationAmount float64 `json:"precipitationAmount"`

	// (Required) The chance of precipitation forecasted to occur during the period.
	PrecipitationChance float64 `json:"precipitationChance"`

	// (Required) The type of precipitation forecasted to occur during the period.
	PrecipitationType PrecipitationType `json:"precipitationType,omitempty"`

	// (Required) The depth of snow as ice crystals forecasted to occur during the period, in millimeters.
	SnowfallAmount float64 `json:"snowfallAmount"`

	// The direction the wind is forecasted to come from during the period, in degrees.
	WindDirection float64 `json:"windDirection"`

	// (Required) The average speed the wind is forecasted to be during the period, in kilometers per hour.
	WindSpeed float64 `json:"windSpeed"`
}

// NextHourForecast is a minute-by-minute forecast for the next hour.
type NextHourForecast struct {
	ProductData
	NextHourForecastData
}

// NextHourForecastData is the next hour forecast information.
type NextHourForecastData struct {
	// The time the forecast ends.
	ForecastEnd *time.Time `json:"forecastEnd,omitempty"`

	// The time the forecast starts.
	ForecastStart *time.Time `json:"forecastStart,omitempty"`

	// (Required) An array of the forecast minutes.
	Minutes []ForecastMinute `json:"minutes,omitempty"`

	// (Required) An array of the forecast summaries.
	Summary []ForecastPeriodSummary `json:"summary,omitempty"`
}

// ForecastMinute is the precipitation forecast for a specified minute.
type ForecastMinute struct {
	// (Required) The probability of precipitation during this minute.
	PrecipitationChance float64 `json:"precipitationChance"`

	// (Required) The precipitation intensity in millimeters per hour.
	PrecipitationIntensity float64 `json:"precipitationIntensity"`

	// (Required) The start time of the minute.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// ForecastPeriodSummary is the summary for a specified period in the minute forecast.
type ForecastPeriodSummary struct {
	// (Required) The type of precipitation forecasted.
	Condition PrecipitationType `json:"condition,omitempty"`

	// The end time of the forecast.
	EndTime *time.Time `json:"endTime,omitempty"`

	// (Required) The probability of precipitation during this period.
	PrecipitationChance float64 `json:"precipitationChance"`

	// (Required) The precipitation intensity in millimeters per hour.
	PrecipitationIntensity float64 `json:"precipitationIntensity"`

	// (Required) The start time of the forecast.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// CurrentWeather is the current weather conditions for the specified location.
type CurrentWeather struct {
	ProductData
	CurrentWeatherData
}

// CurrentWeatherData is the current weather object.
type CurrentWeatherData struct {
	// (Required) The date and time.
	AsOf *time.Time `json:"asOf,omitempty"`

	// The percentage of the sky covered with clouds during the period, from 0 to 1.
	CloudCover float64 `json:"cloudCover"`

	// (Required) An enumeration value indicating the condition at the time.
	ConditionCode string `json:"conditionCode,omitempty"`

	// A Boolean value indicating whether there is daylight.
	DayLight bool `json:"daylight"`

	// (Required) The relative humidity, from 0 to 1.
	Humidity float64 `json:"humidity"`

	// (Required) The precipitation intensity, in millimeters per hour.
	PrecipitationIntensity float64 `json:"precipitationIntensity"`

	// (Required) The sea level air pressure, in millibars.
	Pressure float64 `json:"pressure"`

	// (Required) The direction of change of the sea-level air pressure.
	PressureTrend PressureTrend `json:"pressureTrend,omitempty"`

	// (Required) The current temperature, in degrees Celsius.
	Temperature float64 `json:"temperature"`

	// (Required) The feels-like temperature when factoring wind and humidity, in degrees Celsius.
	TemperatureApparent float64 `json:"temperatureApparent"`

	// (Required) The temperature at which relative humidity is 100%, in Celsius.
	TemperatureDewPoint float64 `json:"temperatureDewPoint"`

	// (Required) The level of ultraviolet radiation.
	UvIndex int64 `json:"uvIndex"`

	// (Required) The distance at which terrain is visible, in meters.
	Visibility float64 `json:"visibility"`

	// The direction of the wind, in degrees.
	WindDirection float64 `json:"windDirection"`

	// The maximum wind gust speed, in kilometers per hour.
	WindGust float64 `json:"windGust"`

	// (Required) The wind speed, in kilometers per hour.
	WindSpeed float64 `json:"windSpeed"`
}

// WeatherAlertCollection is a collecton of weather alerts.
type WeatherAlertCollection struct {
	// (Required) An array of weather alert summaries.
	Alerts []WeatherAlertSummary `json:"alerts,omitempty"`

	// A URL that provides more information about the alerts.
	DetailsURL string `json:"detailsUrl,omitempty"`
}

// ResponseType is the recommended action from a reporting agency.
type ResponseType string

const (
	// Take shelter in place.
	ResponseTypeShelter ResponseType = "shelter"

	// Relocate.
	ResponseTypeEvacuate ResponseType = "evacuate"

	// Make preparations.
	ResponseTypePrepare ResponseType = "prepare"

	// Execute a pre-planned activity.
	ResponseTypeExecute ResponseType = "execute"

	// Avoid the event.
	ResponseTypeAvoid ResponseType = "avoid"

	// Monitor the situation.
	ResponseTypeMonitor ResponseType = "monitor"

	// Assess the situation.
	ResponseTypeAssess ResponseType = "assess"

	// The event no longer poses a threat.
	ResponseTypeAllClear ResponseType = "allClear"

	// No action recommended.
	ResponseTypeNone ResponseType = "none"
)

// Severity is the level of danger to life and property.
type Severity string

const (
	// Extraordinary threat.
	SeverityExtreme Severity = "extreme"

	// Significant threat.
	SeveritySevere Severity = "severe"

	// Possible threat.
	SeverityModerate Severity = "moderate"

	// Minimal or no known threat.
	SeverityMinor Severity = "minor"

	// Unknown threat.
	SeverityUnknown Severity = "unknown"
)

// Urgency is an indication of urgency of action from the reporting agency.
type Urgency string

const (
	// Take responsive action immediately.
	UrgencyImmediate Urgency = "immediate"

	// Take responsive action in the next hour.
	UrgencyExpected Urgency = "expected"

	// Take responsive action in the near future.
	UrgencyFuture Urgency = "future"

	// Responsive action is no longer required.
	UrgencyPast Urgency = "past"

	// The urgency is unknown.
	UrgencyUnknown Urgency = "unknown"
)

// Certainty is how likely the event is to occur.
type Certainty string

const (
	// The event has already occurred or is ongoing.
	CertaintyObserved Certainty = "observed"

	// The event is likely to occur (greater than 50% probability).
	CertaintyLikely Certainty = "likely"

	// The event is unlikley to occur (less than 50% probability).
	CertaintyPossible Certainty = "possible"

	// The event is not expected to occur (approximately 0% probability).
	CertaintyUnlikely Certainty = "unlikely"

	// It is unknown if the event will occur.
	CertaintyUnknown Certainty = "unknown"
)

// WeatherAlertSummary contains detailed information about the weather alert.
type WeatherAlertSummary struct {
	// An official designation of the affected area.
	AreaID string `json:"areaId,omitempty"`

	// A human-readable name of the affected area.
	AreaName string `json:"areaName,omitempty"`

	// (Required) How likely the event is to occur.
	Certainty Certainty `json:"certainty,omitempty"`

	// (Required) The ISO code of the reporting country.
	CountryCode string `json:"countryCode,omitempty"`

	// (Required) A human-readable description of the event.
	Description string `json:"description,omitempty"`

	// The URL to a page containing detailed information about the event.
	DetailsURL string `json:"detailsUrl,omitempty"`

	// (Required) The time the event went into effect.
	EffectiveTime *time.Time `json:"effectiveTime,omitempty"`

	// The time when the underlying weather event is projected to end.
	EventEndTime *time.Time `json:"eventEndTime,omitempty"`

	// The time when the underlying weather event is projected to start.
	EventOnSetTime *time.Time `json:"eventOnSetTime,omitempty"`

	// (Required) The time when the event expires.
	ExpireTime *time.Time `json:"expireTime,omitempty"`

	// (Required) A unique identifier of the event.
	ID string `json:"id,omitempty"`

	// (Required) The time that event was issued by the reporting agency.
	IssuedTime *time.Time `json:"issuedTime,omitempty"`

	// (Required) An array of recommended actions from the reporting agency.
	Responses []ResponseType `json:"responses,omitempty"`

	// (Required) The level of danger to life and property.
	Severity Severity `json:"severity,omitempty"`

	// (Required) The name of the reporting agency.
	Source string `json:"source,omitempty"`

	// An indication of urgency of action from the reporting agency.
	Urgency Urgency `json:"urgency,omitempty"`
}

// ProductData is a base type for all weather data.
type ProductData struct {
	// The name of the data set.
	Name DataSet `json:"name,omitempty" yaml:"name,omitempty"`

	// (Required) Descriptive information about the weather data.
	Metadata Metadata `json:"metadata,omitempty"`
}

// UnitsSystem is the system of units that the weather data is reported in.
type UnitsSystem string

const (
	// The metric system.
	UnitsMetric UnitsSystem = "m"
)

// Metadata holds descriptive information about the weather data.
type Metadata struct {
	// The URL of the legal attribution for the data source.
	AttributionURL string `json:"attributionURL,omitempty"`

	// (Required) The time when the weather data is no longer valid.
	ExpireTime *time.Time `json:"expireTime,omitempty"`

	// The ISO language code for localizable fields.
	Language string `json:"language,omitempty"`

	// (Required) The latitude of the relevant location.
	Latitude float64 `json:"latitude"`

	// (Required) The longitude of the relevant location.
	Longitude float64 `json:"longitude"`

	// The URL of a logo for the data provider.
	ProviderLogo string `json:"providerLogo,omitempty"`

	// The name of the data provider.
	ProviderName string `json:"providerName,omitempty"`

	// (Required) The time the weather data was procured.
	ReadTime *time.Time `json:"readTime,omitempty"`

	// The time the provider reported the weather data.
	ReportedTime *time.Time `json:"reportedTime,omitempty"`

	// The weather data is temporarily unavailable from the provider.
	TemporarilyUnavailable bool `json:"temporarilyUnavailable,omitempty"`

	// The system of units that the weather data is reported in.
	Units UnitsSystem `json:"units,omitempty"`

	// (Required) The data format version.
	Version int `json:"version"`
}
