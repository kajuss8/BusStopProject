package models

type ServiceAvailability uint8

const (
	ServiceAdded   ServiceAvailability = 1
	ServiceRemoved ServiceAvailability = 2
)

type CalendarDates struct {
	Id            int              `json:"id"`
	ServiceId     int              `json:"serviceId"`
	Date          string           `json:"date"`
	ExceptionType ServiceAvailability `json:"exceptionType"`
}