package models

type StopTime struct {
	TripID        string `gorm:"not null;index" json:"tripId"` 
	ArrivalTime   string `gorm:"not null" json:"arrivalTime"`
	DepartureTime string `gorm:"not null" json:"departureTime"`
	StopID        int `gorm:"not null;index" json:"stopId"` 
	StopSequence  int    `gorm:"not null" json:"stopSequence"`
}