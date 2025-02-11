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

const routeFileName = "routes.txt"

func getAllRoutes() ([]Route, error) {
	var routesResult []Route
	routes, err := handleFiles.ReadFile(filepath + routeFileName)
	if err != nil {
		return nil, fmt.Errorf("GetAllRoutes failed: %w", err)
	}

	for _, route := range routes {
		routeId := route[0]
		routeShortName := route[1]
		routeLongName := route[2]
		routeDesc := route[3]
		routeType, err := strconv.Atoi(route[4])
		if err != nil {
			return nil, fmt.Errorf("getAllRoutes failed to parse routeType: %w", err)
		}
		routeUrl := route[5]
		routeColor := route[6]
		routeTextColor := route[7]
		routeSortOrder, err := strconv.Atoi(route[8])
		if err != nil {
			return nil, fmt.Errorf("getAllRoutes failed to parse routeSortOrder: %w", err)
		}

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

func convertTripIdToRoutesShortLongNameAndType(routeIds []string) (shortName, longName, routeT []string, err error) {
	routes, err := getAllRoutes()
	if err != nil{
		return nil, nil, nil, err
	}

	routeMap := make(map[string]Route, len(routes))
	for _, route := range routes {
		routeMap[route.RouteId] = route
	}

	var routeType []TransportType
	for _, routeId := range routeIds {
		if route, exists := routeMap[routeId]; exists {
			shortName = append(shortName, route.RouteShortName)
			longName = append(longName, route.RouteLongName)
			routeType = append(routeType, route.RouteType)
		}
	}
    return shortName, longName, convertRouteTypeNumberToLetter(routeType), nil
}

// func getRouteLongNameById(routeId string) (string, error) {
// 	routes, err := getAllRoutes()
// 	if err != nil {
// 		return "", err
// 	}

// 	routeMap := make(map[string]string, len(routes))
// 	for _, route := range routes {
// 		routeMap[route.RouteId] = route.RouteLongName
// 	}

// 	longName, exists := routeMap[routeId]
// 	if !exists {
// 		return "", fmt.Errorf("route with ID %s not found", routeId)
// 	}

// 	return longName, nil
// }

func convertRouteTypeNumberToLetter(routeTypes []TransportType) (routeLetter []string) {
	for _, routeType := range routeTypes{
		switch routeType {
		case Bus:
			routeLetter = append(routeLetter, "A")
		case Trolleybus:
			routeLetter = append(routeLetter, "T")
		}
	}
	return routeLetter
}