package models

import (
	"busProject/src/handleFiles"
	"fmt"
	"strconv"
)

type PickupStatus int

const (
	RegularPickup    PickupStatus = 0
	NoPickup         PickupStatus = 1
	PhoneAgency      PickupStatus = 2
	CoordinateDriver PickupStatus = 3
)

type DropOffStatus int

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

func getAllStopTimes() (stopTimesResult []StopTime, err error){
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

func getStopTimesByStopId(stopId string) (stopTimesById []StopTime, err error) {
	stopTimes, err := getAllStopTimes()
	if err != nil {
		return nil, err
	}

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

func getStopTimesByTripIds(tripIds [][]string) (stopTimesByTripIds [][]StopTime, err error) {
	allStopTimes, err := getAllStopTimes()
	if err != nil {
		return nil, err
	}

	tripIdMap := make(map[string][]StopTime)
	for _, stopTime := range allStopTimes {
		tripIdMap[stopTime.TripId] = append(tripIdMap[stopTime.TripId], stopTime)
	}

	for _, tripIdGroup := range tripIds {
		var stopTimesGroup []StopTime
		for _, tripId := range tripIdGroup {
			if stopTimes, exists := tripIdMap[tripId]; exists {
				stopTimesGroup = append(stopTimesGroup, stopTimes...)
			}
		}
		stopTimesByTripIds = append(stopTimesByTripIds, stopTimesGroup)
	}

	return stopTimesByTripIds, nil
}

func orderStopTimesBySequence(stopTimes [][]StopTime) [][]StopTime {
	for i, stopTimesSlice := range stopTimes {
		sequenceMap := make(map[int][]StopTime)
		for _, stopTime := range stopTimesSlice {
			sequenceMap[stopTime.StopSequence] = append(sequenceMap[stopTime.StopSequence], stopTime)
		}

		var groupedStopTimes []StopTime
		for seq := 0; seq <= len(sequenceMap); seq++ {
			if stops, exists := sequenceMap[seq]; exists {
				groupedStopTimes = append(groupedStopTimes, stops...)
			}
		}
		stopTimes[i] = groupedStopTimes
	}
	return stopTimes
}

func getDepartureTimesByStopIds(stopTimes [][]StopTime) ([][][]string, error) {
	var result [][][]string

	for _, stopTimesSlice := range stopTimes {
		stopIdMap := make(map[string][]string)
		for _, stopTime := range stopTimesSlice {
			stopIdMap[stopTime.StopId] = append(stopIdMap[stopTime.StopId], stopTime.DepartureTime[:5])
		}

		var groupedDepartureTimes [][]string
		processedStopIds := make(map[string]bool)
		for _, stopTime := range stopTimesSlice {
			if _, exists := processedStopIds[stopTime.StopId]; !exists {
				groupedDepartureTimes = append(groupedDepartureTimes, stopIdMap[stopTime.StopId])
				processedStopIds[stopTime.StopId] = true
			}
		}
		result = append(result, groupedDepartureTimes)
	}

	return result, nil
}

func getUniqueStopIds(stopTimes [][]StopTime) [][]string {
	var uniqueStopIds [][]string

	for _, stopTimesSlice := range stopTimes {
		stopIdMap := make(map[string]bool)
		var stopIds []string
		for _, stopTime := range stopTimesSlice {
			if _, exists := stopIdMap[stopTime.StopId]; !exists {
				stopIdMap[stopTime.StopId] = true
				stopIds = append(stopIds, stopTime.StopId)
			}
		}
		uniqueStopIds = append(uniqueStopIds, stopIds)
	}

	return uniqueStopIds
}

func getTripIdsFromStopTime(stopTimes []StopTime) (tripIds []string) {
	for _, stopTime := range stopTimes {
		tripIds = append(tripIds, stopTime.TripId)
	}
	return tripIds
}

func convertTripIdToStopTimesArrivalTime(mappedTrip [][]Trip, stopTimesByStopId []StopTime) (arrivalTimes [][]string, err error) {
	stopTimeMap := make(map[string]string, len(stopTimesByStopId))
	for _, stopTime := range stopTimesByStopId {
		stopTimeMap[stopTime.TripId] = stopTime.ArrivalTime
	}

	for _, trips := range mappedTrip {
		var times []string
		for _, trip := range trips {
			if arrivalTime, exists := stopTimeMap[trip.TripId]; exists {
				times = append(times, arrivalTime[:5])
			} else {
				return nil, fmt.Errorf("ConvertTripIdToStopTimesArrivalTime failed: no such trip ID")
			}
		}
		arrivalTimes = append(arrivalTimes, times)
	}
    return arrivalTimes, nil
}