package models

import (
	"busProject/src/handleFiles"
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
	TripId        string	`json:"tripId"`
	ArrivalTime   string	`json:"arrivalTime"`
	DepartureTime string	`json:"departureTime"`
	StopId        int	`json:"stopId"`
	StopSequence  int	`json:"stopSequence"`
	PickupType    PickupStatus	`json:"pickupType"`
	DropOffType   DropOffStatus	`json:"dropOffType"`
}

const stopTimeFilePath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/gtfsFolder/stop_times.txt"
const stopIdColumnName = "stop_id"

func GetAllArriveTimesByStopId(stopId string) ([]StopTime, error) {

	var stopTimes []StopTime
	arriveTimes, err := handleFiles.ReadFromCsvAllWithSameId(stopTimeFilePath, stopIdColumnName, stopId)
	if err != nil {
		return nil, err
	}

	for _, arriveTime := range arriveTimes {
		tripId := arriveTime[0]
		arrivalTime := arriveTime[1]
		departureTime := arriveTime[2]
		stopId, _ := strconv.Atoi(arriveTime[3])
		stopSequence, _ := strconv.Atoi(arriveTime[4])
		pickupType, _ := strconv.Atoi(arriveTime[5])
		dropOffType, _ := strconv.Atoi(arriveTime[6])

		stopTimes = append(stopTimes ,StopTime{
			TripId:        tripId,
			ArrivalTime:   arrivalTime,
			DepartureTime: departureTime,
			StopId:        stopId,
			StopSequence:  stopSequence,
			PickupType:    PickupStatus(pickupType),
			DropOffType:   DropOffStatus(dropOffType),
		})
	}
	return stopTimes, nil
}