package models

type RouteSchedule struct {
	RouteLongName string
	ShapeId 	string
	RouteInfo	[]RouteInformation
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
	workDays, err := convertServiceIdToCalendarDays(serviceIds)
	if err != nil {
		return nil, err
	}

	stopTimes, err := GetStopTimesByTripIds(tripIds)
	if err != nil {
		return nil, err
	}
	stopTimes = OrderStopTimesBySequence(stopTimes)

	stopIds := GetUniqueStopIds(stopTimes)
	stopNames, err := GetStopNames(stopIds)
	if err != nil {
		return nil, err
	}

	routeSchedules, err := buildRouteSchedules(shapeIds, workDays, stopNames, routeLongName)
	if err != nil {
		return nil, err
	}

	return routeSchedules, nil
}

func buildRouteSchedules(shapeIds []string, workDays [][]int, stopNames [][]string, routeLongName []string) ([]RouteSchedule, error) {
	var routeSchedules []RouteSchedule
	for i, shapeId := range shapeIds {
		stopInfo := buildStopInfo(stopNames[i])
		routeSchedules = append(routeSchedules, RouteSchedule{
			RouteLongName: routeLongName[i],
			ShapeId: shapeId,
			RouteInfo: []RouteInformation{
				{
					WorkDays: workDays[i],
					StopInfo: stopInfo,
				},
			},
		})
	}
	return routeSchedules, nil
}

func buildStopInfo(stopNames []string) []StopInfo {
	var stopInfo []StopInfo
	for _, stopName := range stopNames {
		stopInfo = append(stopInfo, StopInfo{
			StopName: stopName,
		})
	}
	return stopInfo
}

func createDifferentRouteLongName(trips [][]Trip, routeId string) ([]string, error) {
	tripHeadsign, direction := getTripHeadsignAndDirection(trips)
	routeName, err := getRouteLongNameById(routeId)
	if err != nil{
		return nil, err
	}

	var routeNames []string
	for range trips {
		routeNames = append(routeNames, routeName)
	}
	
	return createRouteLongName(routeNames, tripHeadsign, direction), nil
}