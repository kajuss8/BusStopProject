package models

import (
	"busProject/src/handleFiles"
	"fmt"
	"strconv"
)

type DayServiceAvailability int

const (
	ServiceAvailable    DayServiceAvailability = 1
	ServiceNotAvailable DayServiceAvailability = 0
)

type Calendar struct {
	ServiceId       int			`json:"serviceId"`
	WeekDaysService []int		`json:"weekServices"`
	StartDate       string		`json:"startDate"`
	EndDate         string		`json:"endDate"`
}

const CalendarFileName = "calendar.txt"

func GetAllCalendars() ([]Calendar, error) {
	var calendarsResult []Calendar
	calendars, err := handleFiles.ReadFile(filepath + CalendarFileName)
	if err != nil {
		return nil, fmt.Errorf("GetAllCalendars failed: %w", err)
	}

	for _, calendar := range calendars {
		serviceId, _ := strconv.Atoi(calendar[0])
		monday, _ := strconv.Atoi(calendar[1])
		tuesday, _ := strconv.Atoi(calendar[2])
		wednesday, _ := strconv.Atoi(calendar[3])
		thursday, _ := strconv.Atoi(calendar[4])
		friday, _ := strconv.Atoi(calendar[5])
		saturday, _ := strconv.Atoi(calendar[6])
		sunday, _ := strconv.Atoi(calendar[7])
		startDate := calendar[8]
		endDate := calendar[9]

		calendarsResult = append(calendarsResult, Calendar{
			ServiceId: serviceId,
			WeekDaysService: []int{
				monday,
				tuesday,
				wednesday,
				thursday,
				friday,
				saturday,
				sunday},
			StartDate: startDate,
			EndDate:   endDate,
		})
	}
	return calendarsResult, nil
}

func ConvertServiceIdToCalendarDays(serviceIds []int) ([][]int, error){
	calendars, err := GetAllCalendars()
	if err != nil {
		return nil, err
	}

	calendarMap := make(map[int][]int, len(calendars))
	for _, calendar := range calendars {
		calendarMap[calendar.ServiceId] = calendar.WeekDaysService
	}

	var result [][]int
	for _, serviceId := range serviceIds {
		if days, exists := calendarMap[serviceId]; exists {
			result = append(result, days)
		} else {
			return nil, fmt.Errorf("ConvertServiceIdToCalendarDays failes: no such service ID")
		}
	}
	return result, nil
}