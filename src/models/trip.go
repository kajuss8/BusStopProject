package models


type Direction uint8

const (
	Outbound Direction = 0
	Inbound  Direction = 1
)

type WheelchairAccessibility uint8

const (
	NoInfo        WheelchairAccessibility = 0
	Accessible    WheelchairAccessibility = 1
	NotAccessible WheelchairAccessibility = 2
)

type Trip struct {
	Id                   int	`json:"id"`
	RouteId              string	`json:"routeId"`
	ServiceId            int	`json:"serviceId"`
	TripId               string	`json:"tripId"`
	TripHeadsign         string	`json:"tripHeadsign"`
	DirectionId          Direction	`json:"directionId"`
	BlockId              int	`json:"blockId"`
	ShapeId              string	`json:"shapeId"`
	WheelchairAccessible WheelchairAccessibility	`json:"wheelchairAccessible"`
}