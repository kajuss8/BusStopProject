package models

import (
	"strings"
	"time"
)

type StopSchedule struct {
	StopName         string            `json:"stopName"`
	StopInformations []StopInformation `json:"stopInformation"`
}

type StopInformation struct {
	ShapeId				string	 `json:"shapeId"`
	RouteId				string	 `json:"routeId"`
	RouteShortName 		string 	 `json:"routeShortName"`
	RouteLongName  		string 	 `json:"routeLongName"`
	RouteType			string 	 `json:"routeType"`
	CalendarWorkDays 	[]string    `json:"workDays"`
	ArrivalTime      	[]string `json:"arrivalTimes"`
}

func CreateStopsSchedule(stopId int) (StopSchedule, error) {
	stop, err := createStop(stopId)
	if err != nil {
		return StopSchedule{}, err
	}

	stopTimes, err := createStopTimes(stopId)
	if err != nil {
		return StopSchedule{}, err
	}

	tripIds := takeStopTimeTripIds(stopTimes)
	trips, err := createTripsByStopTimeTripIds(tripIds)
	if err != nil {
		return StopSchedule{}, err
	}

	mappedTripShape, shapeId := tripsShapeIdMapped(trips)

	tripRouteIds := getTripsShapeRouteId(mappedTripShape)
	serviceIds := getTripsShapeServiceIds(mappedTripShape)

	sName, lName, routeType, err := convertTripIdToRoutesShortLongNameAndType(tripRouteIds)
	if err != nil {
		return StopSchedule{}, err
	}

	calWorkDays, startDate, endDate, err := convertServiceIdToCalendarDays(serviceIds)
	if err != nil {
		return StopSchedule{}, err
	}
	namedWorkDays := createWorkDaysLetterSchedule(calWorkDays)

	arrivalTimes, err := convertTripIdToStopTimesArrivalTime(mappedTripShape, stopTimes)
	if err != nil {
		return StopSchedule{}, err
	}

	tripHeadsign, direction := getTripHeadsignAndDirection(mappedTripShape)
	routeLongName := createRouteLongName(lName, tripHeadsign, direction)

	currentDate := time.Now()
	stopSchedule := assembleStopSchedule(stop, mappedTripShape, sName, routeLongName, routeType, namedWorkDays, arrivalTimes, startDate, endDate, currentDate, shapeId, tripRouteIds)

	return stopSchedule, nil
}

func assembleStopSchedule(stop Stop, mappedTripShape [][]Trip, sName, routeLongName, routeType []string, calWorkDays [][]string, arrivalTimes [][]string,
	 startDate, endDate []time.Time, currentDate time.Time, shapeId, routeId []string) StopSchedule {
	stopSchedule := StopSchedule{
		StopName: getStopName(stop),
	}
	
	for i := 0; i < len(mappedTripShape); i++ {
		if currentDate.After(startDate[i]) && currentDate.Before(endDate[i]) {
			info := StopInformation{
				ShapeId: shapeId[i],
				RouteId: routeId[i],
				RouteShortName:   sName[i],
				RouteLongName:    routeLongName[i],
				RouteType:        routeType[i],
				CalendarWorkDays: calWorkDays[i],
				ArrivalTime:      arrivalTimes[i],
			}
			stopSchedule.StopInformations = append(stopSchedule.StopInformations, info)
		}
	}

	return stopSchedule
}

func createStop(stopId int) (Stop, error) {
	stop, err := getStopById(stopId)
	if err != nil {
		return stop, err
	}

	return stop, nil
}

func createStopTimes(stopId int) ([]StopTime, error) {
	stopTimes, err := getStopTimesByStopId(stopId)
	if err != nil {
		return nil, err
	}
	return stopTimes, nil
}

func takeStopTimeTripIds(stopTimes []StopTime) (tripIds []string) {
	tripIds = getTripIdsFromStopTime(stopTimes)
	return tripIds
}

func createTrips() ([]Trip, error) {
	return getAllTrips()
}

func createTripsByStopTimeTripIds(tripIds []string) ([]Trip, error) {
	Trips, err := createTrips()
	if err != nil {
		return nil, err
	}

	tripsByIds, err := getTripsByTripIds(tripIds, Trips)
	if err != nil{
		return nil, err
	}
	return tripsByIds, nil
}

func createRouteLongName(lName []string, tHeadSign []string, direction []int) (routeLongName []string) {
	for i, name := range lName {
		parts := strings.Split(name, " - ")
		if direction[i] == 1 {
			for j, k := 0, len(parts)-1; j < k; j, k = j+1, k-1 {
				parts[j], parts[k] = parts[k], parts[j]
			}
		}
		if parts[len(parts)-1] != tHeadSign[i] {
			parts[len(parts)-1] = tHeadSign[i]
		}
		routeLongName = append(routeLongName, strings.Join(parts, " - "))
	}
	return routeLongName
}

func createWorkDaysLetterSchedule(workDays [][]DayServiceAvailability) (result [][]string) {
	for _, workday := range workDays {
		result = append(result, convertCalendarDaysToLetters(workday))
	}
	return result
}

