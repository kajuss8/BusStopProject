package models

type FareRule struct {
	Id 			uint32 	`json:"id"`
	FareId 		uint32 	`json:"fareId"`
	RouteId 	string 	`json:"routeId"`
	OriginId 	uint16 	`json:"originId"`
	DestinationId uint16 `json:"destinationId"`
	ContainsId 	uint16 	`json:"containsId"`
}