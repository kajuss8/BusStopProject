package models

type TransportType uint8

const (
    TramStreetcarLightRail TransportType = 0
    SubwayMetro            TransportType = 1
    Rail                   TransportType = 2
    Bus                    TransportType = 3
    Ferry                  TransportType = 4
    CableTram              TransportType = 5
    AerialLift             TransportType = 6
    Funicular              TransportType = 7
    Trolleybus             TransportType = 11
    Monorail               TransportType = 12
)

type Route struct {
	Id 				uint32 			`json:"id"`
	RouteId 		string 			`json:"routeId"`
	RouteShortName 	string 			`json:"routeShortName"`
	RouteLongName 	string 			`json:"routeLongName"`
	RouteDesc 		string 			`json:"routeDesc"`
	RouteType 		TransportType 	`json:"routeType"`
	RouteUrl 		string 			`json:"routeUrl"`
	RouteColor 		string 			`json:"routeColor"`
	RouteTextColor 	string 			`json:"routeTextColor"`
	RouteSortOrder 	uint32 			`json:"routeSortOrder"`
}