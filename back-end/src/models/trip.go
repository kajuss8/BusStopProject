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

func getAllTrips() ([]Trip, error) {
	var TripsResult []Trip
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

func getTripsByIds(StopTimetripIds []string, trips []Trip) ([]Trip, error) {
	tripMap := make(map[string]Trip, len(StopTimetripIds))
	for _, trip := range trips {
		tripMap[trip.TripId] = trip
	}

	var routeIdsResult []Trip
	for _, tripId := range StopTimetripIds {
		if trip, exists := tripMap[tripId]; exists {
			routeIdsResult = append(routeIdsResult, trip)
		} else {
			return nil, fmt.Errorf("GetTripsByIds failed: no such trip ID")
		}
	}
	return routeIdsResult, nil
}

func tripsShapeIdMapped(trips []Trip) [][]Trip {
	var groupedTrips [][]Trip
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
			groupedTrips = append(groupedTrips, []Trip{trip})
		}
	}
	return groupedTrips
}

func getTripHeadsignAndDirection(trips [][]Trip) (tHeadSign []string, tDirection []int){
	var headsign []string
	var direction []int
	for _, trip := range trips{
		direction = append(direction, int(trip[0].DirectionId))
		headsign = append(headsign, trip[0].TripHeadsign)
	}
	return headsign, direction
}

func getTripsShapeServiceIds(shapeIdsMapped [][]Trip) []int {
	var result []int
	for _, value := range shapeIdsMapped{
		for _, v := range value {
			result = append(result, v.ServiceId)
			break
		}
	}
	return result
}

func getTripsShapeRouteId(shapeIdsMapped [][]Trip) []string {
	var result []string
	for _, value := range shapeIdsMapped{
		for _, v := range value {
			result = append(result, v.RouteId)
			break
		}
	}
	return result
}