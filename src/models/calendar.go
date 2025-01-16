package models

import "time"

type DayServiceAvailability uint8

const (
	ServiceAvailable    DayServiceAvailability = 1
	ServiceNotAvailable DayServiceAvailability = 0
)

type Calendar struct {
	Id              uint32                            `json:"id"`
	ServiceId       uint32                            `json:"serviceId"`
	WeekDaysService map[string]DayServiceAvailability `json:"weekServices"`
	StartDate       time.Time                         `json:"startDate"`
	EndDate         time.Time                         `json:"endDate"`
}