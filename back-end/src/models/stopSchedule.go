package models

import "strings"

type StopSchedule struct {
	StopName         string            `json:"stopName"`
	StopInformations []StopInformation `json:"stopInformation"`
}

type StopInformation struct {
	RouteShortName 		string 	 `json:"routeShortName"`
	RouteLongName  		string 	 `json:"routeLongName"`
	RouteType			string 	 `json:"routeType"`
	CalendarWorkDays 	[]int    `json:"workDays"`
	ArrivalTime      	[]string `json:"arrivalTimes"`
}

func CreateStopsSchedule(stopId string) (StopSchedule, error) {
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
	mappedTripShape := tripsShapeIdMapped(trips)

	tripRouteIds := getTripsShapeRouteId(mappedTripShape)
	serviceIds := getTripsShapeServiceIds(mappedTripShape)

	sName, lName, routeType, err := convertTripIdToRoutesShortLongNameAndType(tripRouteIds)
	if err != nil{
		return StopSchedule{}, err
	}
	calWorkDays, err := convertServiceIdToCalendarDays(serviceIds)
	if err != nil{
		return StopSchedule{}, err
	}
	arrivalTimes, err := ConvertTripIdToStopTimesArrivalTime(mappedTripShape, stopTimes)
	if err != nil{
		return StopSchedule{}, err
	}

	tripHeadsign, direction := getTripHeadsignAndDirection(mappedTripShape)
	routeLongName := createRouteLongName(lName, tripHeadsign, direction)

	var stopSchedule StopSchedule
	stopSchedule.StopName = getStopName(stop)
	for i := 0; i < len(mappedTripShape); i++ {
		info := struct {
			RouteShortName   string   `json:"routeShortName"`
			RouteLongName    string   `json:"routeLongName"`
			RouteType		 string 	`json:"routeType"`
			CalendarWorkDays []int    `json:"workDays"`
			ArrivalTime      []string `json:"arrivalTimes"`
		}{
			RouteLongName:    routeLongName[i],
			RouteShortName:   sName[i],
			RouteType: 		routeType[i],
			CalendarWorkDays: calWorkDays[i],
			ArrivalTime:      arrivalTimes[i],
		}
		stopSchedule.StopInformations = append(stopSchedule.StopInformations, info)
	}
	return stopSchedule, nil
}

func createStop(stopId string) (Stop, error) {
	stop, err := getStopById(stopId)
	if err != nil {
		return stop, err
	}

	return stop, nil
}

func createStopTimes(stopId string) ([]StopTime, error) {
	stopTimes, err := GetStopTimesByStopId(stopId)
	if err != nil {
		return nil, err
	}
	return stopTimes, nil
}

func takeStopTimeTripIds(stopTimes []StopTime) (tripIds []string) {
	trips := GetTripIds(stopTimes)
	return trips
}

func createTrips() ([]Trip, error) {
	return getAllTrips()
}

func createTripsByStopTimeTripIds(tripIds []string) ([]Trip, error) {
	Trips, err := createTrips()
	if err != nil {
		return nil, err
	}

	tripsByIds, err := getTripsByIds(tripIds, Trips)
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