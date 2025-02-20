package models

type Stop struct {
	StopID   int  `gorm:"primaryKey" json:"stopId"` 
	StopName string  `gorm:"not null" json:"stopName"`
	StopLat  float64 `gorm:"not null" json:"stopLat"`
	StopLon  float64 `gorm:"not null" json:"stopLon"`
	StopURL  string  `gorm:"not null" json:"stopUrl"`
}