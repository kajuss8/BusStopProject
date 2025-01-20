package models

import (
	"busProject/src/handleFiles"
	"errors"
	"strconv"
	"strings"
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

func GetAllStopTimes() ([]StopTime, error) {
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

func GetStopTimesByStopId(stopId string) ([]StopTime, error){
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

func GetArrivalTimes(stopId string) ([]string, error) {
	stopTimes, err := GetStopTimesByStopId(stopId)
	if err != nil {
		return nil, err
	}

	var arrivalTimes []string
	for _, stopTime := range stopTimes {
		arrivalTimes = append(arrivalTimes, stopTime.ArrivalTime)
	}

	return arrivalTimes, nil
}

func GetTripId(stopId string) ([]string, error) {
	stopTimes, err := GetStopTimesByStopId(stopId)
	if err != nil {
		return nil, err
	}

	if len(stopTimes) == 0 {
		return nil, errors.New("Stop not found")
	}

	var tripIds []string
	for _, stopTime := range stopTimes {
		tripIds = append(tripIds, stopTime.TripId)
	}

	return tripIds, nil
}

type ArriveTime struct {
	BusType string
	BusNumber string
	ArrivalTime []string
}

func GetBusTypeAndNumber(tripId string) (busType string, busNumber string) {
	index := strings.Index(tripId, "-")
	if index != -1 {
		substring := tripId[:index]
		busType := string(substring[0])
		busNumber := substring[1:]
		return busType, busNumber
	}
	return "", ""
}

func GetArriveTimesById(stopId string) ([]ArriveTime, error) {
	stopTimesById, err := GetStopTimesByStopId(stopId)
	if err != nil {
		return nil, err
	}

	arriveTimeMap := make(map[string]ArriveTime)

	for _, stopTime := range stopTimesById {
		busType, busNumber := GetBusTypeAndNumber(stopTime.TripId)
		key := busType + busNumber

		if at, exists := arriveTimeMap[key]; exists {
			at.ArrivalTime = append(at.ArrivalTime, stopTime.ArrivalTime)
			arriveTimeMap[key] = at
		} else {
			arriveTimeMap[key] = ArriveTime{
				BusType:    busType,
				BusNumber:  busNumber,
				ArrivalTime: []string{stopTime.ArrivalTime},
			}
		}
	}

	var arrivalTime []ArriveTime
	for _, at := range arriveTimeMap {
		arrivalTime = append(arrivalTime, at)
	}

	return arrivalTime, nil
}