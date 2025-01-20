package models

import (
	"busProject/src/handleFiles"
	"errors"
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

const stopTimeFilePath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/gtfsFolder/stop_times.txt"

func GetAllStopTimes() ([]StopTime, error){
	var stopTimesResult []StopTime
	stopTimes, err := handleFiles.ReadFile(stopTimeFilePath)
	if err != nil {
		return nil, err
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
		return nil, errors.New("Stop not found")
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

func GetSingleTripId(stopTimes []StopTime) string {
	return stopTimes[0].TripId
}

func GetSequence(stopTimes []StopTime) []int {
	var sequences []int
	for _, stopTime := range stopTimes {
		sequences = append(sequences, stopTime.StopSequence)
	}

	return sequences
}