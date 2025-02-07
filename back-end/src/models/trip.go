package models

import (
	"busProject/src/handleFiles"
	"fmt"
	"strconv"
)

type Direction int

const (
	Outbound Direction = 0
	Inbound  Direction = 1
)

type WheelchairAccessibility uint8

const (
	NoInfo        WheelchairAccessibility = 0
	Accessible    WheelchairAccessibility = 1
	NotAccessible WheelchairAccessibility = 2
)

type Trip struct {
	RouteId              string                  `json:"routeId"`
	ServiceId            int                     `json:"serviceId"`
	TripId               string                  `json:"tripId"`
	TripHeadsign         string                  `json:"tripHeadsign"`
	DirectionId          Direction               `json:"directionId"`
	BlockId              int                     `json:"blockId"`
	ShapeId              string                  `json:"shapeId"`
	WheelchairAccessible WheelchairAccessibility `json:"wheelchairAccessible"`
}

const tripFileName = "trips.txt"

func getAllTrips() (TripsResult []Trip, err error) {
	trips, err := handleFiles.ReadFile(filepath + tripFileName)
	if err != nil {
		return nil, fmt.Errorf("GetAllTrips failed: %w", err)
	}

	for _, trip := range trips {
		routeId := trip[0]
		serviceId, _ := strconv.Atoi(trip[1])
		tripId := trip[2]
		tripHeadsign := trip[3]
		directionId, _ := strconv.Atoi(trip[4])
		blockId, _ := strconv.Atoi(trip[5])
		shapeId := trip[6]
		wheelchairAccessible, _ := strconv.Atoi(trip[7])

		TripsResult = append(TripsResult, Trip{
			RouteId:              routeId,
			ServiceId:            serviceId,
			TripId:               tripId,
			TripHeadsign:         tripHeadsign,
			DirectionId:          Direction(directionId),
			BlockId:              blockId,
			ShapeId:              shapeId,
			WheelchairAccessible: WheelchairAccessibility(wheelchairAccessible),
		})
	}
	return TripsResult, nil
}

func getTripsByRouteId(routeId string) (tripsResult []Trip, err error) {
	trips, err := getAllTrips()
	if err != nil {
		return nil, err
	}

	for _, trip := range trips {
		if trip.RouteId == routeId {
			tripsResult = append(tripsResult, trip)
		}
	}
	if(tripsResult == nil){
		return tripsResult, fmt.Errorf("getTripsByRouteId failed: no such trip ID")
	}

	return tripsResult, nil
}

func getTripIds(trips [][]Trip) (tripIds [][]string) {
	for _, tripGroup := range trips {
		var groupTripIds []string
		for _, trip := range tripGroup {
			groupTripIds = append(groupTripIds, trip.TripId)
		}
		tripIds = append(tripIds, groupTripIds)
	}
	return tripIds
}

func getTripsByTripIds(StopTimetripIds []string, trips []Trip) (routeIds []Trip, err error) {
	tripMap := make(map[string]Trip, len(StopTimetripIds))
	for _, trip := range trips {
		tripMap[trip.TripId] = trip
	}

	for _, tripId := range StopTimetripIds {
		if trip, exists := tripMap[tripId]; exists {
			routeIds = append(routeIds, trip)
		} else {
			return nil, fmt.Errorf("GetTripsByIds failed: no such trip ID")
		}
	}
	return routeIds, nil
}

func tripsShapeIdMapped(trips []Trip) (groupedTrips [][]Trip, shapeIds []string) {
	for _, trip := range trips {
		found := false
		for i := range groupedTrips {
			if groupedTrips[i][0].ShapeId == trip.ShapeId && groupedTrips[i][0].ServiceId == trip.ServiceId {
				groupedTrips[i] = append(groupedTrips[i], trip)
				found = true
				break
			}
		}
		if !found {
			shapeIds = append(shapeIds, trip.ShapeId)
			groupedTrips = append(groupedTrips, []Trip{trip})
		}
	}
	return groupedTrips, shapeIds
}

func getTripHeadsignAndDirection(trips [][]Trip) (headsign []string, direction []int){
	for _, trip := range trips{
		direction = append(direction, int(trip[0].DirectionId))
		headsign = append(headsign, trip[0].TripHeadsign)
	}
	return headsign, direction
}

func getTripsShapeServiceIds(shapeIdsMapped [][]Trip) (serviceIds []int) {
	for _, trip := range shapeIdsMapped{
		serviceIds = append(serviceIds, trip[0].ServiceId)
	}
	return serviceIds
}

func getTripsShapeRouteId(shapeIdsMapped [][]Trip) (routeIds []string) {
	for _, trip := range shapeIdsMapped{
		routeIds = append(routeIds, trip[0].RouteId)
	}
	return routeIds
}