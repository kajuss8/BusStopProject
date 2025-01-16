package models

import "time"

type PickupStatus uint8

const (
	RegularPickup    PickupStatus = 0
	NoPickup         PickupStatus = 1
	PhoneAgency      PickupStatus = 2
	CoordinateDriver PickupStatus = 3
)

type DropOffStatus uint8

const (
	RegularDropOff          DropOffStatus = 0
	NoDropOff               DropOffStatus = 1
	PhoneAgencyDropOff      DropOffStatus = 2
	CoordinateDriverDropOff DropOffStatus = 3
)

type StopTime struct {
	Id            uint32        `json:"id"`
	TripId        string        `json:"tripId"`
	ArrivalTime   time.Time     `json:"arrivalTime"`
	DepartureTime time.Time     `json:"departureTime"`
	StopId        uint16        `json:"stopId"`
	StopSequence  uint8         `json:"stopSequence"`
	PickupType    PickupStatus  `json:"pickupType"`
	DropOffType   DropOffStatus `json:"dropOffType"`
}