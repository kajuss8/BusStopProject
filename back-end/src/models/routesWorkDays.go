package models

import (
	"busProject/internal/gtfs/handleFiles"
	"fmt"
	"strconv"
)

type RouteWorkDays struct {
	RouteId            string   `json:"routeId"`
	RouteShortName     string   `json:"routeShortName"`
	RouteLongName      string   `json:"routeLongName"`
	RouteTransportType string   `json:"routeTransportType"`
	CalendarWorkDays   []string `json:"workDays"`
}

func CreateRouteWorkDays() ([]RouteWorkDays, error) {
	routes, err := GetAllRoutes()
	if err != nil {
		return nil, err
	}

	diffTripServiceId, err := GetTripServiceIds()
	if err != nil {
		return nil, err
	}
	updated := UpdateServiceIds(diffTripServiceId)
	merged := MergeServiceDays(updated)


	

	return assembleRouteWorkDays(routes, merged), nil
}

func assembleRouteWorkDays(routes []Route, mergedServiceDays map[string][]string) []RouteWorkDays {
	var routeWorkDays []RouteWorkDays

	for _, route := range routes {
		if serviceDays, exist := mergedServiceDays[route.RouteId]; exist {
			routeWorkDays = append(routeWorkDays, RouteWorkDays{
				RouteId: route.RouteId,
				RouteShortName: route.RouteShortName,
				RouteLongName: route.RouteLongName,
				RouteTransportType: route.RouteTransportType,
				CalendarWorkDays: serviceDays,
			})
		}
	}
	return routeWorkDays
}

func GetTripServiceIds() (TripRouteIdMapped map[string][]int, err error) {
	trips, err := handleFiles.ReadFile(filepath + tripFileName)
	if err != nil {
		return nil, fmt.Errorf("GetAllTrips failed: %w", err)
	}

	m := make(map[string][]int)

	for _, trip := range trips {
		routeId := trip[0]
		serviceId, err := strconv.Atoi(trip[1])
		if err != nil {
			return nil, fmt.Errorf("getTripServiceIds failed to parse serviceId: %w", err)
		}

		
		if !contains(m[routeId], serviceId) {
			m[routeId] = append(m[routeId], serviceId)
		}
	}

	return m, nil
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func UpdateServiceIds(tripMap map[string][]int) map[string][][]DayServiceAvailability {
	calendars, err := getAllCalendars()
	if err != nil {
		return nil
	}

	calendarMap := make(map[int][]DayServiceAvailability)
	for _, cal := range calendars {
		calendarMap[cal.ServiceId] = cal.WeekDaysService
	}

	updatedMap := make(map[string][][]DayServiceAvailability)

	for routeId, serviceIds := range tripMap {
		var updatedServiceIds [][]DayServiceAvailability

		for _, serviceId := range serviceIds {
			if days, exists := calendarMap[serviceId]; exists {
				updatedServiceIds = append(updatedServiceIds, days)
			}
		}

		updatedMap[routeId] = updatedServiceIds
	}

	return updatedMap
}

func MergeServiceDays(updatedMap map[string][][]DayServiceAvailability) map[string][]string {
	mergedMap := make(map[string][]string)

	for routeId, serviceSlices := range updatedMap {
		if len(serviceSlices) == 0 {
			continue
		}

		mergedDays := make([]DayServiceAvailability, len(serviceSlices[0]))
		copy(mergedDays, serviceSlices[0])

		for _, serviceDays := range serviceSlices[1:] {
			for i := range serviceDays {
				if mergedDays[i] == 0 && serviceDays[i] == 1 {
					mergedDays[i] = 1
				}
			}
		}

		mergedMap[routeId] = convertCalendarDaysToLetters(mergedDays)
	}

	return mergedMap
}