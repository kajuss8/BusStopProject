package database

import (
	"fmt"
	"busProject/internal/gtfs/handleFiles"
	"strconv"
	"gorm.io/gorm"
	"busProject/models"
)


func LoadStops(db *gorm.DB, filePath, stopFileName string) error {
	stops, err := handleFiles.ReadFile(filePath + stopFileName)
	if err != nil {
		return fmt.Errorf("LoadStops failed: %w", err)
	}
	
	for _, stop := range stops {
		stopId, err := strconv.Atoi(stop[0])
		if err != nil {
			return fmt.Errorf("LoadStops failed to parse stopId: %w", err)
		}
		stopName := stop[2]
		stopLat, err := strconv.ParseFloat(stop[4], 64)
		if err != nil {
			return fmt.Errorf("LoadStops failed to parse stopLat: %w", err)
		}
		stopLon, err := strconv.ParseFloat(stop[5], 64)
		if err != nil {
			return fmt.Errorf("LoadStops failed to parse stopLon: %w", err)
		}
		stopUrl := stop[6]

		stop := models.Stop{
			StopID:        stopId,
			StopName:      stopName,
			StopLat:       stopLat,
			StopLon:       stopLon,
			StopURL:       stopUrl,
		}

		if err := db.Create(&stop).Error; err != nil {
			return fmt.Errorf("LoadStops failes: %w", err)
		}
	}
	fmt.Println("Stops loaded successfully!")
	return nil
}

func LoadCalendar(db *gorm.DB, filePath, calendarFileName string) error {
	calendar, err := handleFiles.ReadFile(filePath + calendarFileName)
	if err != nil {
		return fmt.Errorf("LoadCalendar failed: %w", err)
	}
	
	for _, cal := range calendar {
		serviceId, err := strconv.Atoi(cal[0])
		if err != nil {
			return fmt.Errorf("LoadCalendar failed to parse serviceId: %w", err)
		}
		monday, err := strconv.Atoi(cal[1])
		if err != nil {
			return fmt.Errorf("LoadCalendar failed to parse monday: %w", err)
		}
		tuesday, err := strconv.Atoi(cal[2])
		if err != nil {
			return fmt.Errorf("LoadCalendar failed to parse tuesday: %w", err)
		}
		wednesday, err := strconv.Atoi(cal[3])
		if err != nil {
			return fmt.Errorf("LoadCalendar failed to parse wednesday: %w", err)
		}
		thursday, err := strconv.Atoi(cal[4])
		if err != nil {
			return fmt.Errorf("LoadCalendar failed to parse thursday: %w", err)
		}
		friday, err := strconv.Atoi(cal[5])
		if err != nil {
			return fmt.Errorf("LoadCalendar failed to parse friday: %w", err)
		}
		saturday, err := strconv.Atoi(cal[6])
		if err != nil {
			return fmt.Errorf("LoadCalendar failed to parse saturday: %w", err)
		}
		sunday, err := strconv.Atoi(cal[7])
		if err != nil {
			return fmt.Errorf("LoadCalendar failed to parse sunday: %w", err)
		}
		startDate := cal[8]
		endDate := cal[9]

		calendar := models.Calendar{
			ServiceID: serviceId,
			Monday:    monday,
			Tuesday:   tuesday,
			Wednesday: wednesday,
			Thursday:  thursday,
			Friday:    friday,
			Saturday:  saturday,
			Sunday:    sunday,
			StartDate: startDate,
			EndDate:   endDate,
		}

		if err := db.Create(&calendar).Error; err != nil {
			return fmt.Errorf("LoadCalendar failes: %w", err)
		}
	}
	fmt.Println("Calendar loaded successfully!")
	return nil
}

func LoadRoutes(db *gorm.DB, filePath, routeFileName string) error {
	routes, err := handleFiles.ReadFile(filePath + routeFileName)
	if err != nil {
		return fmt.Errorf("LoadRoutes failed: %w", err)
	}

	for _, route := range routes {
		routeId := route[0]
		routeShortName := route[1]
		routeLongName := route[2]
		routeType, err := strconv.Atoi(route[4])
		if err != nil {
			return fmt.Errorf("LoadRoutes failed to parse routeType: %w", err)
		}

		routeUrl := route[5]
		routeColor := route[6]
		routeTextColor := route[7]
		routeSortOrder, err := strconv.Atoi(route[8])
		if err != nil {
			return fmt.Errorf("LoadRoutes failed to parse routeSortOrder: %w", err)
		}

		route := models.Route{
			RouteID:        routeId,
			RouteShortName: routeShortName,
			RouteLongName:  routeLongName,
			RouteType:      routeType,
			RouteURL:       routeUrl,
			RouteColor:     routeColor,
			RouteTextColor: routeTextColor,
			RouteSortOrder: routeSortOrder,
		}

		if err := db.Create(&route).Error; err != nil {
			return fmt.Errorf("LoadRoutes failes: %w", err)
		}
	}
	return nil
}

func LoadTrips(db *gorm.DB, filePath, tripFileName string) error {
	trips, err := handleFiles.ReadFile(filePath + tripFileName)
	if err != nil {
		return fmt.Errorf("LoadTrips failed: %w", err)
	}

	for _, trip := range trips {
		routeId := trip[0]
		serviceId, err := strconv.Atoi(trip[1])
		if err != nil {
			return fmt.Errorf("LoadTrips failed to parse serviceId: %w", err)
		}
		tripId := trip[2]
		tripHeadsign := trip[3]
		directionId, err := strconv.Atoi(trip[4])
		if err != nil {
			return fmt.Errorf("LoadTrips failed to parse directionId: %w", err)
		}
		blockId, err := strconv.Atoi(trip[5])
		if err != nil {
			return fmt.Errorf("LoadTrips failed to parse blockId: %w", err)
		}
		shapeId := trip[6]

		trip := models.Trip{
			TripID:       tripId,
			RouteID:      routeId,
			TripHeadsign: tripHeadsign,
			DirectionID:  directionId,
			BlockID:      blockId,
			ShapeID:      shapeId,
			ServiceID:    serviceId,
		}

		if err := db.Create(&trip).Error; err != nil {
			return fmt.Errorf("LoadTrips failes: %w", err)
		}
	}
	fmt.Println("Trips loaded successfully!")
	return nil
}

func LoadStopTimes(db *gorm.DB, filePath, stopTimeFileName string) error {
	stopTimes, err := handleFiles.ReadFile(filePath + stopTimeFileName)
	if err != nil {
		return fmt.Errorf("LoadStopTimes failed: %w", err)
	}

	for _, stopTime := range stopTimes {
		tripID := stopTime[0]
		arrivalTime := stopTime[1]
		departureTime := stopTime[2]
		stopID, err := strconv.Atoi(stopTime[3])
		if err != nil {
			return fmt.Errorf("LoadStopTimes failed to parse stopID: %w", err)
		}
		stopSequence, err := strconv.Atoi(stopTime[4])
		if err != nil {
			return fmt.Errorf("LoadStopTimes failed to parse stopSequence: %w", err)
		}

		stopTime := models.StopTime{
			TripID:        tripID,
			ArrivalTime:   arrivalTime,
			DepartureTime: departureTime,
			StopID:        stopID,
			StopSequence:  stopSequence,
		}

		if err := db.Create(&stopTime).Error; err != nil {
			return fmt.Errorf("LoadStopTimes failes: %w", err)
		}
	}

	return nil
}