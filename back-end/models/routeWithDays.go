package models

type RouteWithDays struct {
	RouteID        string   `json:"routeId"`
	RouteShortName string   `json:"routeShortName"`
	RouteLongName  string   `json:"routeLongName"`
	RouteType      int      `json:"routeType"`
	RouteURL       string   `json:"routeUrl"`
	RouteColor     string   `json:"routeColor"`
	RouteTextColor string   `json:"routeTextColor"`
	RouteSortOrder int      `json:"routeSortOrder"`
	WeekDays       []string `json:"weekDays"`
}