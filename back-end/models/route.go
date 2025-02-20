package models

type Route struct {
	RouteID        string `gorm:"primaryKey" json:"routeId"`
	RouteShortName string `gorm:"not null" json:"routeShortName"`
	RouteLongName  string `gorm:"not null" json:"routeLongName"`
	RouteType      int    `gorm:"not null" json:"routeType"`
	RouteURL       string `gorm:"not null" json:"routeUrl"`
	RouteColor     string `gorm:"not null" json:"routeColor"`
	RouteTextColor string `gorm:"not null" json:"routeTextColor"`
	RouteSortOrder int    `gorm:"not null" json:"routeSortOrder"`
}
