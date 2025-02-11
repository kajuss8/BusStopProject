package models

import (
	"sort"
	"time"
)

type RouteSchedule struct {
	RouteLongName string	`json:"routeLongName"`
	ShapeId       string	`json:"shapeId"`
	RouteInfo     []RouteInformation	`json:"routeInfo"`
}

type RouteInformation struct {
	WorkDays []string	`json:"workDays"`
	StopInfo []StopInfo	`json:"stopInfo"`
}

type StopInfo struct {
	StopId		  int	`json:"stopId"`
	StopName      string	`json:"stopName"`
	DepartureTime []string	`json:"departureTime"`
}

func CreateRouteSchedule(routeId string) ([]RouteSchedule, error) {
	trips, err := getTripsByRouteId(routeId)
	if err != nil {
		return nil, err
	}
	tripsmappedByShapeIds, shapeIds := tripsShapeIdMapped(trips)
	tripIds := getTripIds(tripsmappedByShapeIds)

	serviceIds := getTripsShapeServiceIds(tripsmappedByShapeIds)
	workDays, startDates, endDates, err := convertServiceIdToCalendarDays(serviceIds)
	if err != nil {
		return nil, err
	}
	namedWorkDays := createWorkDaysWordsSchedule(workDays)

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

	routeLongName, err := createRoutesLongName(stopIds)
	if err != nil {
		return nil, err
	}

	departureTimes, err := getDepartureTimesByStopIds(stopTimes)
	if err != nil {
		return nil, err
	}

	currentDate := time.Now()
	return buildRouteSchedules(shapeIds, namedWorkDays, stopNames, stopIds, routeLongName, departureTimes, startDates, endDates, currentDate)
}

func buildRouteSchedules(shapeIds []string, workDays [][]string, stopNames [][]string, stopId [][]int, routeLongName []string, departureTimes [][][]string, startDates, endDates []time.Time, currentDate time.Time) ([]RouteSchedule, error) {
	shapeIdMap := make(map[string]*RouteSchedule)
	for i, shapeId := range shapeIds {
		if currentDate.After(startDates[i]) && currentDate.Before(endDates[i]){
				stopInfo := buildStopInfo(stopNames[i], stopId[i], departureTimes[i])
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

func buildStopInfo(stopNames []string, stopId []int, departureTimes [][]string) []StopInfo {
	var stopInfo []StopInfo
	for i, stopName := range stopNames {
		stopInfo = append(stopInfo, StopInfo{
			StopId: 	stopId[i],
			StopName:      stopName,
			DepartureTime: departureTimes[i],
		})
	}
	return stopInfo
}

func createRoutesLongName(stopIds [][]int) (routeLName []string, err error) {
	for _, stopIdSlice := range stopIds{
		firstStopIdName, err := getStopNameById(stopIdSlice[0])
		if err != nil {
			return nil, err
		}
		lastStopIdName, err := getStopNameById(stopIdSlice[len(stopIdSlice)-1])
		if err != nil {
			return nil, err
		}
		temp := firstStopIdName + " - " + lastStopIdName

		routeLName = append(routeLName, temp)
	}

	return routeLName, nil
}

func createWorkDaysWordsSchedule(workDays [][]DayServiceAvailability) (result [][]string) {
	for _, workday := range workDays {
		result = append(result, convertCalendarDaysToWords(workday))
	}
	return result
}
