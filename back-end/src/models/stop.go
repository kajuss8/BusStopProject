package models

import (
	"busProject/src/handleFiles"
	"fmt"
	"strconv"
)

type LocationCategory uint8

const (
	StopOrPlatform LocationCategory = 0
	Station        LocationCategory = 1
	EntranceExit   LocationCategory = 2
	GenericNode    LocationCategory = 3
	BoardingArea   LocationCategory = 4
)

type Stop struct {
	StopId        string           `json:"stopId"`
	StopCode      string           `json:"stopCode"`
	StopName      string           `json:"stopName"`
	StopDesc      string           `json:"stopDesc"`
	StopLat       float64          `json:"stopLat"`
	StopLon       float64          `json:"stopLon"`
	StopUrl       string           `json:"stopUrl"`
	LocationType  LocationCategory `json:"locationType"`
	ParentStation int              `json:"parentStation"`
}

const (
	filepath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/back-end/gtfsFolder/"
	stopFileName = "stops.txt"
)

func GetAllStops() ([]Stop, error) {
	var stopsResult []Stop
	stops, err := handleFiles.ReadFile(filepath + stopFileName)
	if err != nil {
		return nil, fmt.Errorf("GetAllStops failed: %w", err)
	}
	
	for _, stop := range stops {
		stopId := stop[0]
		stopCode := stop[1]
		stopName := stop[2]
		stopDesc := stop[3]
		stopLat, _ := strconv.ParseFloat(stop[4], 64)
		stopLon, _ := strconv.ParseFloat(stop[5], 64)
		stopUrl := stop[6]
		locationType, _ := strconv.Atoi(stop[7])
		parentStation, _ := strconv.Atoi(stop[8])

		stopsResult = append(stopsResult, Stop{
			StopId:        stopId,
			StopCode:      stopCode,
			StopName:      stopName,
			StopDesc:      stopDesc,
			StopLat:       stopLat,
			StopLon:       stopLon,
			StopUrl:       stopUrl,
			LocationType:  LocationCategory(locationType),
			ParentStation: parentStation,
		})
	}
	return stopsResult, nil
}

func GetStopById(stopId string) (Stop, error) {
	stops, err := GetAllStops()
	if err != nil {
		return Stop{}, err
	}

	for _, stop := range stops {
		if stop.StopId == stopId {
			return stop, nil
		}
	}
	return Stop{}, fmt.Errorf("GetStopId failed: no such stop ID")
}

func GetStopName(stop Stop) string {
	return stop.StopName
}
