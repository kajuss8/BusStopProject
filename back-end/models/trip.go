package models

type Trip struct {
	TripID       string `gorm:"primaryKey" json:"tripId"`
	RouteID      string `gorm:"not null;index" json:"routeId"` 
	TripHeadsign string `gorm:"not null" json:"tripHeadsign"`
	DirectionID  int    `gorm:"not null" json:"directionId"`
	BlockID      int    `gorm:"not null" json:"blockId"`
	ShapeID      string `gorm:"not null;index" json:"shapeId"`
	ServiceID    int `gorm:"not null;index" json:"serviceId"` 
}