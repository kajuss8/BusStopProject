package models

type LocationCategory uint8

const (
	StopOrPlatform LocationCategory = 0
	Station        LocationCategory = 1
	EntranceExit   LocationCategory = 2
	GenericNode    LocationCategory = 3
	BoardingArea   LocationCategory = 4
)

type Stop struct {
	StopId        string		`json:"stopId"`
	StopCode      string		`json:"stopCode"`
	StopName      string		`json:"stopName"`
	StopDesc      string		`json:"stopDesc"`
	StopLat       float32		`json:"stopLat"`
	StopLon       float32		`json:"stopLon"`
	StopUrl       string		`json:"stopUrl"`
	LocationType  LocationCategory	`json:"locationType"`
	ParentStation int		    `json:"parentStation"`
}

//const stopFilePath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/gtfsFolder/stop_times.txt"
//const stopIdColumnName = "stop_id"



