package handleBusSchedule

import "busProject/src/models"

type BusSchedule struct {
	StopName		string           		`json:"stopName"`
	RouteShortName 	string					`json:"routeShortName"`
	RouteLongName	string					`json:"routeLongName"`
	DirectionId		models.Direction 		`json:"directionId"`
	RouteType		models.TransportType 	`json:"routeType"`
	ArrivalTimes 	[]string         		`json:"arrivalTimes"`
	CalendarDays 	map[string]models.DayServiceAvailability 	`json:"CalendarWorkDays"` 
}

// func CreateBusSchedule(stopId string) (BusSchedule, error){
// 	var busSchedule BusSchedule
// 	busSchedule.StopName = models.GetStopName(stopId)
// 	busSchedule.RouteShortName, _ = models.GetRouteShortName(getRouteIdByStopId(stopId))
// 	busSchedule.RouteLongName, _ = models.GetRouteLongName(getRouteIdByStopId(stopId))
// 	busSchedule.DirectionId, _ = models.GetDirection(getTripIdByStopId(stopId))
// 	busSchedule.RouteType, _ = models.GetRouteTypeById(getRouteIdByStopId(stopId))
// 	busSchedule.ArrivalTimes, _ = models.GetArrivalTimes(stopId)
// 	busSchedule.CalendarDays, _ = models.GetCalendarWorkDays(getServiceIdByStopId(stopId))

// 	return busSchedule, nil
// }

// func getRouteIdByStopId(stopId string) (string){
// 	tripId := getTripIdByStopId(stopId)
// 	routeId, _ := models.GetRouteId(tripId)
// 	return routeId
// }

// func getTripIdByStopId(stopId string) (string){
// 	tripId, _ := models.GetSingleTripId(stopId)
// 	return tripId
// }

// func getServiceIdByStopId(stopId string) (int){
// 	tripId := getTripIdByStopId(stopId)
// 	serviceId, _ := models.GetServiceId(tripId)
// 	return serviceId
// }