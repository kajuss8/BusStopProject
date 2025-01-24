package models

type FareRule struct {
	FareId 		int 	`json:"fareId"`
	RouteId 	string 	`json:"routeId"`
	OriginId 	int 	`json:"originId"`
	DestinationId int `json:"destinationId"`
	ContainsId 	int 	`json:"containsId"`
}