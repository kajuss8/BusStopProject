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

func GetAllStopTimes() (stopTimesResult []StopTime, err error){
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

func GetStopTimesByStopId(stopId string) (stopTimesById []StopTime, err error) {
	stopTimes, err := GetAllStopTimes()
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

func GetStopTimesByTripIds(tripIds [][]string) (stopTimesByTripIds [][]StopTime, err error) {
	allStopTimes, err := GetAllStopTimes()
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

func OrderStopTimesBySequence(stopTimes [][]StopTime) [][]StopTime {
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

func GetDepartureTimesByStopIds(stopTimes [][]StopTime, stopIds [][]string) ([][]string) {
	return nil
}
// func GetDepartureTimesByStopIds(stopTimes [][]StopTime, stopIds [][]string) ([][]string, error) {
// 	var result [][]string

// 	for _, stopTimesSlice := range stopTimes {
// 		stopIdMap := make(map[string][]string)
// 		for _, stopTime := range stopTimesSlice {
// 			stopIdMap[stopTime.StopId] = append(stopIdMap[stopTime.StopId], stopTime.DepartureTime[:5])
// 		}

// 		for _, stopIdGroup := range stopIds {
// 			var departureTimes []string
// 			for _, stopId := range stopIdGroup {
// 				if times, exists := stopIdMap[stopId]; exists {
// 					departureTimes = append(departureTimes, times...)
// 				} else {
// 					return nil, fmt.Errorf("GetDepartureTimesByStopIds failed: no such stop ID")
// 				}
// 			}
// 			result = append(result, departureTimes)
// 		}
// 	}

// 	return result, nil
// }

func OrderStopTimesByDepartureTime(stopTimes [][]StopTime) [][]StopTime {
	for i, stopTimesSlice := range stopTimes {
		timeMap := make(map[string][]StopTime)
		for _, stopTime := range stopTimesSlice {
			timeMap[stopTime.DepartureTime] = append(timeMap[stopTime.DepartureTime], stopTime)
		}

		var groupedStopTimes []StopTime
		for _, stopTime := range stopTimesSlice {
			if stops, exists := timeMap[stopTime.DepartureTime]; exists {
				groupedStopTimes = append(groupedStopTimes, stops...)
				delete(timeMap, stopTime.DepartureTime)
			}
		}
		stopTimes[i] = groupedStopTimes
	}
	return stopTimes
}

func GetUniqueStopIds(stopTimes [][]StopTime) [][]string {
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

func GetArrivalTimes(stopTimes []StopTime)  (arrivalTimes []string) {
	for _, stopTime := range stopTimes {
		arrivalTimes = append(arrivalTimes, stopTime.ArrivalTime[:5])
	}
	return arrivalTimes
}

func GetTripIds(stopTimes []StopTime) (tripIds []string) {
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

func GetAllDepartureTimes(stopTimesByTripIds [][]StopTime) (allDepartureTimes [][]string) {
	for _, stopTimes := range stopTimesByTripIds {
		var DepartureTimes []string
		for _, stopTime := range stopTimes {
			DepartureTimes = append(DepartureTimes, stopTime.DepartureTime[:5])
		}
		allDepartureTimes = append(allDepartureTimes, DepartureTimes)
	}
	return allDepartureTimes
}

func GetAllStopIds(stopTimesByTripIds [][]StopTime) (allStopIds [][]string) {
	for _, stopTimes := range stopTimesByTripIds {
		var stopIds []string
		for _, stopTime := range stopTimes {
			stopIds = append(stopIds, stopTime.StopId)
		}
		allStopIds = append(allStopIds, stopIds)
	}
	return allStopIds
}