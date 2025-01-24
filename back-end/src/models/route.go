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

func GetAllRoutes() ([]Route, error) {
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
		}else{
			return nil, fmt.Errorf("GetDifferentRouts failed: no such route ID")
		}
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

func ConvertTripIdToRoutesShortLongNameAndType(routeIds []string) (shortName []string, longName []string, routeT []string, err error) {
	routes, err := GetAllRoutes()
	if err != nil{
		return nil, nil, nil, err
	}

	routeMap := make(map[string]Route, len(routes))
	for _, route := range routes {
		routeMap[route.RouteId] = route
	}

	var sName []string
	var lName []string
	var routeType []TransportType
	for _, routeId := range routeIds {
		if route, exists := routeMap[routeId]; exists {
			sName = append(sName, route.RouteShortName)
			lName = append(lName, route.RouteLongName)
			routeType = append(routeType, route.RouteType)
		}
	}
    return sName, lName, ConvertRouteTypeNumberToLetter(routeType), nil
}

func ConvertRouteTypeNumberToLetter(routeTypes []TransportType) []string {
	var routeLetter []string
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