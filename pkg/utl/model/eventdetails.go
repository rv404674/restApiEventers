package restApiEventers

type EventDetails struct {
	Base
	EventName string `json:"event_name"`
	EventLocation string `json:"event_location"`
	EventDate string `json:"event_date"`
}