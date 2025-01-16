package models

import "time"

type ServiceAvailability uint8

const (
	ServiceAdded   ServiceAvailability = 1
	ServiceRemoved ServiceAvailability = 2
)

type CalendarDates struct {
	Id            uint32              `json:"id"`
	ServiceId     uint32              `json:"serviceId"`
	Date          time.Time           `json:"date"`
	ExceptionType ServiceAvailability `json:"exceptionType"`
}