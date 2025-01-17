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
	Id            uint16           `json:"id"`
	StopId        string           `json:"stopId"`
	StopCode      string           `json:"stopCode"`
	StopName      string           `json:"stopName"`
	StopDesc      string           `json:"stopDesc"`
	StopLat       float32          `json:"stopLat"`
	StopLon       float32          `json:"stopLon"`
	StopUrl       string           `json:"stopUrl"`
	LocationType  LocationCategory `json:"locationType"`
	ParentStation int			   `json:"parentStation"`
}

func CreateStops() []Stop {
	return []Stop{
		{1, "1", "1", "Stop 1", "Stop 1 Description", 1.0, 1.0, "http://stop1.com", StopOrPlatform, 0},
		{2, "2", "2", "Stop 2", "Stop 2 Description", 2.0, 2.0, "http://stop2.com", StopOrPlatform, 0},
		{3, "3", "3", "Stop 3", "Stop 3 Description", 3.0, 3.0, "http://stop3.com", EntranceExit, 0},
		{4, "4", "4", "Stop 4", "Stop 4 Description", 4.0, 4.0, "http://stop4.com", StopOrPlatform, 0},
		{5, "5", "5", "Stop 5", "Stop 5 Description", 5.0, 5.0, "http://stop5.com", Station, 0},
	}
}
