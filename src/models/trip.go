package models

import (
	"busProject/src/handleFiles"
	"errors"
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

const tripFilePath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/gtfsFolder/trips.txt"

func GetAllTrips() ([]Trip, error) {
	var TripsResult []Trip
	trips, err := handleFiles.ReadFile(tripFilePath)
	if err != nil {
		return nil, err
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

	if len(TripsResult) == 0 {
		return nil, errors.New("Stop not found")
	}
	return TripsResult, nil
}

func GetTripById(tripId string) (Trip, error) {
	trips, err := GetAllTrips()
	if err != nil {
		return Trip{}, err
	}

	for _, trip := range trips {
		if trip.TripId == tripId {
			return trip, nil
		}
	}
	return Trip{}, errors.New("Trip not found")
}

func GetRouteId(tripId string) (string, error) {
	trip, err := GetTripById(tripId)
	if err != nil {
		return "", err
	}

	return trip.RouteId, nil
}

func GetServiceId(tripId string) (int, error) {
	trip, err := GetTripById(tripId)
	if err != nil {
		return 0, err
	}

	return trip.ServiceId, nil
}

func GetTripHeadsign(tripId string) (string, error){
	trip, err := GetTripById(tripId)
	if err != nil {
		return "", err
	}

	return trip.TripHeadsign, nil
}

func GetDirection(tripId string) (Direction, error){
	trip, err := GetTripById(tripId)
	if err != nil {
		return 0, err
	}

	return trip.DirectionId, err
}