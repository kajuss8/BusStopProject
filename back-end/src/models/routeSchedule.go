package models

import (
	"sort"
	"time"
)

type RouteSchedule struct {
	RouteLongName string
	ShapeId       string
	RouteInfo     []RouteInformation
}

type RouteInformation struct {
	WorkDays []int
	StopInfo []StopInfo
}

type StopInfo struct {
	StopName      string
	DepartureTime []string
}

func CreateRouteSchedule(routeId string) ([]RouteSchedule, error) {
	trips, err := getTripsByRouteId(routeId)
	if err != nil {
		return nil, err
	}
	tripsmappedByShapeIds, shapeIds := tripsShapeIdMapped(trips)
	tripIds := getTripIds(tripsmappedByShapeIds)

	routeLongName, err := createDifferentRouteLongName(tripsmappedByShapeIds, routeId)
	if err != nil {
		return nil, err
	}

	serviceIds := getTripsShapeServiceIds(tripsmappedByShapeIds)
	workDays, startDates, endDates, err := convertServiceIdToCalendarDays(serviceIds)
	if err != nil {
		return nil, err
	}

	stopTimes, err := getStopTimesByTripIds(tripIds)
	if err != nil {
		return nil, err
	}
	stopTimes = orderStopTimesBySequence(stopTimes)

	stopIds := getUniqueStopIds(stopTimes)
	stopNames, err := getStopNames(stopIds)
	if err != nil {
		return nil, err
	}

	departureTimes, err := getDepartureTimesByStopIds(stopTimes)
	if err != nil {
		return nil, err
	}

	currentDate := time.Now()
	return buildRouteSchedules(shapeIds, workDays, stopNames, routeLongName, departureTimes, startDates, endDates, currentDate)
}

func buildRouteSchedules(shapeIds []string, workDays [][]int, stopNames [][]string, routeLongName []string, departureTimes [][][]string, startDates, endDates []time.Time, currentDate time.Time) ([]RouteSchedule, error) {
	shapeIdMap := make(map[string]*RouteSchedule)
	for i, shapeId := range shapeIds {
		if currentDate.After(startDates[i]) && currentDate.Before(endDates[i]){
				stopInfo := buildStopInfo(stopNames[i], departureTimes[i])
				routeInfo := RouteInformation{
				WorkDays: workDays[i],
				StopInfo: stopInfo,
			}
			if routeSchedule, exists := shapeIdMap[shapeId]; exists {
				routeSchedule.RouteInfo = append(routeSchedule.RouteInfo, routeInfo)
			} else {
					shapeIdMap[shapeId] = &RouteSchedule{
					RouteLongName: routeLongName[i],
					ShapeId:       shapeId,
					RouteInfo:     []RouteInformation{routeInfo},
				}
			}
		}
	}

	routeSchedules := make([]RouteSchedule, 0, len(shapeIdMap))
	for _, routeSchedule := range shapeIdMap {
		routeSchedules = append(routeSchedules, *routeSchedule)
	}

	sort.Slice(routeSchedules, func(i, j int) bool {
		return routeSchedules[i].ShapeId < routeSchedules[j].ShapeId
	})

	return routeSchedules, nil
}

func buildStopInfo(stopNames []string, departureTimes [][]string) []StopInfo {
	var stopInfo []StopInfo
	for i, stopName := range stopNames {
		stopInfo = append(stopInfo, StopInfo{
			StopName:      stopName,
			DepartureTime: departureTimes[i],
		})
	}
	return stopInfo
}

func createDifferentRouteLongName(trips [][]Trip, routeId string) ([]string, error) {
	tripHeadsign, direction := getTripHeadsignAndDirection(trips)
	routeName, err := getRouteLongNameById(routeId)
	if err != nil {
		return nil, err
	}

	var routeNames []string
	for range trips {
		routeNames = append(routeNames, routeName)
	}

	return createRouteLongName(routeNames, tripHeadsign, direction), nil
}