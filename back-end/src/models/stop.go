package models

import (
	"busProject/src/handleFiles"
	"fmt"
	"strconv"
)

type LocationCategory int

const (
	StopOrPlatform LocationCategory = 0
	Station        LocationCategory = 1
	EntranceExit   LocationCategory = 2
	GenericNode    LocationCategory = 3
	BoardingArea   LocationCategory = 4
)

type Stop struct {
	StopId        int           	`json:"stopId"`
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

func getAllStops() (stopsResult []Stop, err error) {
	stops, err := handleFiles.ReadFile(filepath + stopFileName)
	if err != nil {
		return nil, fmt.Errorf("GetAllStops failed: %w", err)
	}
	
	for _, stop := range stops {
		stopId, err := strconv.Atoi(stop[0])
		if err != nil {
			return nil, fmt.Errorf("getAllStops failed to parse stopId: %w", err)
		}
		stopCode := stop[1]
		stopName := stop[2]
		stopDesc := stop[3]
		stopLat, err := strconv.ParseFloat(stop[4], 64)
		if err != nil {
			return nil, fmt.Errorf("getAllStops failed to parse stopLat: %w", err)
		}
		stopLon, err := strconv.ParseFloat(stop[5], 64)
		if err != nil {
			return nil, fmt.Errorf("getAllStops failed to parse stopLon: %w", err)
		}
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

func getStopById(stopId int) (Stop, error) {
	stops, err := getAllStops()
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

func getStopNameById(stopId int) (string, error) {
	stops, err := getAllStops()
	if err != nil {
		return "", err
	}

	for _, stop := range stops {
		if stop.StopId == stopId {
			return stop.StopName, nil
		}
	}
	return "", fmt.Errorf("getStopNameById failed: no such stop ID")
}

func getStopName(stop Stop) string {
	return stop.StopName
}

func getStopNames(allStopIds [][]int) (stopNames [][]string, err error) {
	stops, err := getAllStops()
	if err != nil {
		return nil, err
	}

	stopMap := make(map[int]string)
	for _, stop := range stops {
		stopMap[stop.StopId] = stop.StopName
	}

	for _, stopIdArr := range allStopIds {
		var stopNameArr []string
		for _, stopId := range stopIdArr {
			if stopName, exists := stopMap[stopId]; exists {
				stopNameArr = append(stopNameArr, stopName)
			} else {
				return nil, fmt.Errorf("getStopNames failed: no such stop ID %v", stopId)
			}
		}
		stopNames = append(stopNames, stopNameArr)
	}
	return stopNames, nil
}

func getStopLanAndLon(allStopIds [][]int) (stopLat, stopLon [][]float64, err error){
	stops, err := getAllStops()
	if err != nil {
		return nil, nil, err
	}

	stopLatMap := make(map[int]float64)
	stopLonMap := make(map[int]float64)

	for _, stop := range stops {
		stopLatMap[stop.StopId] = stop.StopLat
		stopLonMap[stop.StopId] = stop.StopLon
	}

	for _, stopIdArr := range allStopIds {
		var stopLatArr []float64
		var stopLonArr []float64
		for _, stopId := range stopIdArr {
			if stopLat, exists := stopLatMap[stopId]; exists {
				stopLatArr = append(stopLatArr, stopLat)
			} else {
				return nil, nil, fmt.Errorf("getStopLanAndLon failed: no such stop ID %v", stopId)
			}

			if stopLon, exists := stopLonMap[stopId]; exists{
				stopLonArr = append(stopLonArr, stopLon)
			}else {
				return nil, nil, fmt.Errorf("getStopLanAndLon failed: no such stop ID %v", stopId)
			}
		}
		stopLat = append(stopLat, stopLatArr)
		stopLon = append(stopLat, stopLonArr)
	}
	return stopLat, stopLon, nil

}