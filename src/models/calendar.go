package models

import (
	"busProject/src/handleFiles"
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

const CalendarFilePath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/gtfsFolder/calendar.txt"

func GetAllCalendars() ([]Calendar, error) {
	var calendarsResult []Calendar
	calendars, err := handleFiles.ReadFile(CalendarFilePath)
	if err != nil {
		return nil, err
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

func ConvertServiceIdToCalendarDays(serviceIds []int) [][]int{
	calendars, _ := GetAllCalendars()

	var result [][]int
	for _, serviceId := range serviceIds {
		for _, calendar := range calendars {
			if serviceId == calendar.ServiceId{
				result = append(result, calendar.WeekDaysService)
				break
			}
		}
	}
	return result
}

func GetCalendarById(serviceIds []int) (Calendar, error) {
	calendars, err := GetAllCalendars()
	if err != nil {
		return Calendar{}, err
	}

	for _, calendar := range calendars {
		for _, id := range serviceIds {
			if calendar.ServiceId == id {
				return calendar, nil
			}
		}
	}
	return Calendar{}, nil
}