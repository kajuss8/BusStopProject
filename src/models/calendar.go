package models

import (
	"busProject/src/handleFiles"
	"strconv"
)

type DayServiceAvailability uint8

const (
	ServiceAvailable    DayServiceAvailability = 1
	ServiceNotAvailable DayServiceAvailability = 0
)

type Calendar struct {
	ServiceId       int                            `json:"serviceId"`
	WeekDaysService map[string]DayServiceAvailability `json:"weekServices"`
	StartDate       string                         `json:"startDate"`
	EndDate         string                        `json:"endDate"`
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
			WeekDaysService: map[string]DayServiceAvailability{
				"monday":    DayServiceAvailability(monday),
				"tuesday":   DayServiceAvailability(tuesday),
				"wednesday": DayServiceAvailability(wednesday),
				"thursday":  DayServiceAvailability(thursday),
				"friday":    DayServiceAvailability(friday),
				"saturday":  DayServiceAvailability(saturday),
				"sunday":    DayServiceAvailability(sunday),
			},
			StartDate: startDate,
			EndDate:   endDate,
		})
	}
	return calendarsResult, nil
}

func GetCalendarById(serviceId int) (Calendar, error) {
	calendars, err := GetAllCalendars()
	if err != nil {
		return Calendar{}, err
	}

	for _, calendar := range calendars {
		if calendar.ServiceId == serviceId {
			return calendar, nil
		}
	}
	return Calendar{}, nil
}

func GetCalendarWorkDays(calendar Calendar) map[string]DayServiceAvailability {
	return calendar.WeekDaysService
}