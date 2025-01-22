package models

type StopSchedule struct {
	StopName         string            `json:"stopName"`
	StopInformations []StopInformation `json:"stopInformation"`
}

type StopInformation struct {
	RouteShortName string `json:"routeShortName"`
	RouteLongName  string `json:"routeLongName"`
	//RouteType			TransportType `json:"routeType"`
	ArrivalTime []string `json:"arrivalTimes"`
	// CalendarWorkDays	map[string]models.DayServiceAvailability `json:"caleendarWorkDays"`
}

func CreateStopSchedule(stopName string, stopInfo []StopInformation) {

}

func CreateStopInformation(routeShortName string, routeLongName string, routeType string, arrivalTimes []string, calendarWorkDays map[string]DayServiceAvailability) {

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
	mappedTripShape := MapTripsShapeId(trips)

	tripRouteIds, _ := GetMapTripsShapeRouteId(mappedTripShape)
	//maptripId := GetMapTripsShapeTripIds(mappedTripShape)


	sName, lName := ConvertTripIdToRoutesShortAndLongName(tripRouteIds)
	arrivalTimes := ConvertTripIdToStopTimesArrivalTime(mappedTripShape, stopTimes)

	var stopSchedule StopSchedule
	stopSchedule.StopName = GetStopName(stop)
	for i := 0; i < len(mappedTripShape); i++ {
		info := struct {
			RouteShortName string   `json:"routeShortName"`
			RouteLongName  string   `json:"routeLongName"`
			ArrivalTime    []string `json:"arrivalTimes"`
		}{
			RouteLongName:  lName[i],
			RouteShortName: sName[i],
			ArrivalTime:    arrivalTimes[i],
		}
		stopSchedule.StopInformations = append(stopSchedule.StopInformations, info)
	}
	return stopSchedule, nil
}

func createStop(stopId string) (Stop, error) {
	stop, err := GetStopById(stopId)
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
	return GetAllTrips()
}

func createTripsByStopTimeTripIds(tripIds []string) ([]Trip, error) {
	Trips, err := createTrips()
	if err != nil {
		return nil, err
	}

	tripsByIds := GetTripsByIds(tripIds, Trips)
	return tripsByIds, nil
}