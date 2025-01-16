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
	Id            uint16           `json:"Id"`
	StopCode      string           `json:"stopCode"`
	StopName      string           `json:"stopName"`
	StopDesc      string           `json:"stopDesc"`
	StopLat       float32          `json:"stopLat"`
	StopLon       float32          `json:"stopLon"`
	StopUrl       string           `json:"stopUrl"`
	LocationType  LocationCategory `json:"locationType"`
	ParentStation uint16           `json:"parentStation"`
}