package models

import (
	"busProject/src/handleFiles"
	"fmt"
	"strconv"
)

type TransportType int

const (
	TramStreetcarLightRail TransportType = 0
	SubwayMetro            TransportType = 1
	Rail                   TransportType = 2
	Bus                    TransportType = 3
	Ferry                  TransportType = 4
	CableTram              TransportType = 5
	AerialLift             TransportType = 6
	Funicular              TransportType = 7
	Trolleybus             TransportType = 800
	Monorail               TransportType = 12
)

type Route struct {
	RouteId        string        `json:"routeId"`
	RouteShortName string        `json:"routeShortName"`
	RouteLongName  string        `json:"routeLongName"`
	RouteDesc      string        `json:"routeDesc"`
	RouteType      TransportType `json:"routeType"`
	RouteUrl       string        `json:"routeUrl"`
	RouteColor     string        `json:"routeColor"`
	RouteTextColor string        `json:"routeTextColor"`
	RouteSortOrder int           `json:"routeSortOrder"`
}

const routeFilePath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/gtfsFolder/routes.txt"

func GetAllRoutes() ([]Route, error) {
	var routesResult []Route
	routes, err := handleFiles.ReadFile(routeFilePath)
	if err != nil {
		return nil, err
	}

	for _, route := range routes {
		routeId := route[0]
		routeShortName := route[1]
		routeLongName := route[2]
		routeDesc := route[3]
		routeType, _ := strconv.Atoi(route[4])
		routeUrl := route[5]
		routeColor := route[6]
		routeTextColor := route[7]
		routeSortOrder, _ := strconv.Atoi(route[8])

		routesResult = append(routesResult, Route{
			RouteId:        routeId,
			RouteShortName: routeShortName,
			RouteLongName:  routeLongName,
			RouteDesc:      routeDesc,
			RouteType:      TransportType(routeType),
			RouteUrl:       routeUrl,
			RouteColor:     routeColor,
			RouteTextColor: routeTextColor,
			RouteSortOrder: routeSortOrder,
		})
	}

	return routesResult, nil
}

func GetRouteById(routeId string) (Route, error) {
    routes, err := GetAllRoutes()
    if err != nil {
        return Route{}, err
    }

	routeMap := make(map[string]Route, len(routes))
	for _, route := range routes{
		routeMap[route.RouteId] = route
	}
    return routeMap[routeId], nil
}

func GetDifferentRouts(routeIds []string, routes []Route) ([]Route, error) {
	routeMap := make(map[string]Route, len(routes))
	for _, route := range routes {
		routeMap[route.RouteId] = route 
	}

	var routesResult []Route
	for _, routeId := range routeIds{
		if route, exists := routeMap[routeId]; exists{
			routesResult = append(routesResult, route)
		}
	}
	if routesResult == nil{
		return nil, fmt.Errorf("GetDifferentRoutes routesResult is nil")
	}
	
    return routesResult, nil
}

func GetRouteIds(routes []Route) []string {
	var result []string
	for _, route := range routes{
		result = append(result, route.RouteId)
	}

	return result
}

func ConvertTripIdToRoutesShortAndLongName(routeIds []string) (shortName []string, longName []string, err error) {
	routes, err := GetAllRoutes()
	if err != nil{
		return nil, nil, fmt.Errorf("GetAllRoutes failed: %w", err)
	}

	routeMap := make(map[string]Route, len(routes))
	for _, route := range routes {
		routeMap[route.RouteId] = route
	}

	var sName []string
	var lName []string
	for _, routeId := range routeIds {
		if route, exists := routeMap[routeId]; exists {
			sName = append(sName, route.RouteShortName)
			lName = append(lName, route.RouteLongName)
		}
	}
    return sName, lName, nil
}

func GetRouteShortName(route Route) string {
    return route.RouteShortName
}

func GetRouteLongName(route Route) string {
    return route.RouteLongName
}

func GetRouteTypeById(route Route) TransportType {
    return route.RouteType
}