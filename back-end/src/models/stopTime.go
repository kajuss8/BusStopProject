package models

import (
	"busProject/src/handleFiles"
	"fmt"
	"strconv"
)

type PickupStatus uint8

const (
	RegularPickup    PickupStatus = 0
	NoPickup         PickupStatus = 1
	PhoneAgency      PickupStatus = 2
	CoordinateDriver PickupStatus = 3
)

type DropOffStatus uint8

const (
	RegularDropOff          DropOffStatus = 0
	NoDropOff               DropOffStatus = 1
	PhoneAgencyDropOff      DropOffStatus = 2
	CoordinateDriverDropOff DropOffStatus = 3
)

type StopTime struct {
	TripId        string		`json:"tripId"`
	ArrivalTime   string     	`json:"arrivalTime"`
	DepartureTime string     	`json:"departureTime"`
	StopId        string        `json:"stopId"`
	StopSequence  int         	`json:"stopSequence"`
	PickupType    PickupStatus  `json:"pickupType"`
	DropOffType   DropOffStatus `json:"dropOffType"`
}

const stopTimeFileName = "stop_times.txt"

func GetAllStopTimes() ([]StopTime, error){
	var stopTimesResult []StopTime
	stopTimes, err := handleFiles.ReadFile(filepath + stopTimeFileName)
	if err != nil {
		return nil, fmt.Errorf("GetAllStopTimes failed: %w", err)
	}

	for _, stopTime := range stopTimes {
		tripId := stopTime[0]
		arrivalTime := stopTime[1]
		departureTime := stopTime[2]
		stopId := stopTime[3]
		stopSequence, _ := strconv.Atoi(stopTime[4])
		pickupType, _ := strconv.Atoi(stopTime[5])
		dropOffType, _ := strconv.Atoi(stopTime[6])

		stopTimesResult = append(stopTimesResult ,StopTime{
			TripId:        tripId,
			ArrivalTime:   arrivalTime,
			DepartureTime: departureTime,
			StopId:        stopId,
			StopSequence:  stopSequence,
			PickupType:    PickupStatus(pickupType),
			DropOffType:   DropOffStatus(dropOffType),
		})
	}
	return stopTimesResult, nil
}

func GetStopTimesByStopId(stopId string) ([]StopTime, error) {
	stopTimes, err := GetAllStopTimes()
	if err != nil {
		return nil, err
	}

	var stopTimesById []StopTime
	for _, stopTime := range stopTimes {
		if stopTime.StopId == stopId {
			stopTimesById = append(stopTimesById, stopTime)
		}
	}

	if len(stopTimesById) == 0 {
		return nil, fmt.Errorf("GetStopTimesByStopId failed: no such stop ID")
	}
	return stopTimesById, nil
}

func GetArrivalTimes(stopTimes []StopTime) []string {
	var arrivalTimes []string
	for _, stopTime := range stopTimes {
		arrivalTimes = append(arrivalTimes, stopTime.ArrivalTime)
	}
	return arrivalTimes
}

func GetTripIds(stopTimes []StopTime) []string {
	var tripIds []string
	for _, stopTime := range stopTimes {
		tripIds = append(tripIds, stopTime.TripId)
	}
	return tripIds
}

func ConvertTripIdToStopTimesArrivalTime(mappedTrip [][]Trip, stopTimesByStopId []StopTime) (arrivalTimes [][]string, err error) {
	stopTimeMap := make(map[string]string, len(stopTimesByStopId))
	for _, stopTime := range stopTimesByStopId {
		stopTimeMap[stopTime.TripId] = stopTime.ArrivalTime
	}

	var arrTimes [][]string
	for _, trips := range mappedTrip {
		var times []string
		for _, trip := range trips {
			if arrivalTime, exists := stopTimeMap[trip.TripId]; exists {
				times = append(times, arrivalTime)
			} else {
				return nil, fmt.Errorf("ConvertTripIdToStopTimesArrivalTime failed: no such trip ID")
			}
		}
		arrTimes = append(arrTimes, times)
	}
    return arrTimes, nil
}