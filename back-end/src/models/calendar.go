package models

import (
	"busProject/src/handleFiles"
	"fmt"
	"strconv"
	"time"
)

type DayServiceAvailability int

const (
	ServiceAvailable    DayServiceAvailability = 1
	ServiceNotAvailable DayServiceAvailability = 0
)

type Calendar struct {
	ServiceId       int			`json:"serviceId"`
	WeekDaysService []DayServiceAvailability		`json:"weekServices"`
	StartDate       time.Time		`json:"startDate"`
	EndDate         time.Time		`json:"endDate"`
}

const CalendarFileName = "calendar.txt"

func getAllCalendars() (calendarsResult []Calendar, err error) {
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
		startDate, err := time.Parse("20060102", calendar[8])
		if err != nil {
			return nil, fmt.Errorf("GetAllCalendars failed to parse startDate: %w", err)
		}
		endDate, err := time.Parse("20060102", calendar[9])
		if err != nil {
			return nil, fmt.Errorf("GetAllCalendars failed to parse endDate: %w", err)
		}

		calendarsResult = append(calendarsResult, Calendar{
			ServiceId: serviceId,
			WeekDaysService: []DayServiceAvailability{
				DayServiceAvailability(monday),
				DayServiceAvailability(tuesday),
				DayServiceAvailability(wednesday),
				DayServiceAvailability(thursday),
				DayServiceAvailability(friday),
				DayServiceAvailability(saturday),
				DayServiceAvailability(sunday)},
			StartDate: startDate,
			EndDate:   endDate,
		})
	}
	return calendarsResult, nil
}

func convertServiceIdToCalendarDays(serviceIds []int) (calendarDays [][]DayServiceAvailability, startDates, endDates []time.Time, err error) {
	calendars, err := getAllCalendars()
	if err != nil {
		return nil, nil, nil, err
	}

	calendarMap := make(map[int]Calendar, len(calendars))
	for _, calendar := range calendars {
		calendarMap[calendar.ServiceId] = calendar
	}

	for _, serviceId := range serviceIds {
		if calendar, exists := calendarMap[serviceId]; exists {
			calendarDays = append(calendarDays, calendar.WeekDaysService)
			startDates = append(startDates, calendar.StartDate)
			endDates = append(endDates, calendar.EndDate)
		} else {
			return nil, nil, nil, fmt.Errorf("ConvertServiceIdToCalendarDays failed: no such service ID")
		}
	}
	return calendarDays, startDates, endDates, nil
}

func convertCalendarDaysToLetters(workDays []DayServiceAvailability) (result []string) {
	days := []string{"P", "A", "T", "K", "P", "Š", "S"}
	
	for i, workDay := range workDays {
		if workDay == ServiceAvailable{
			result = append(result, days[i])
		}else {
			result = append(result, "0")
		}
	}

	return result
}

func convertCalendarDaysToWords(workDays []DayServiceAvailability) (result []string) {
	days := []string{"pirmadienis", "antradienis", "trečiadienis", "ketvirtadienis", "penktadienis", "šeštadienis", "sekmadienis"}

	for i, workDay := range workDays {
		if workDay == ServiceAvailable{
			result = append(result, days[i])
		}
	}

	return result
}