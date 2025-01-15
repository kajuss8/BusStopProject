package models

import "time"

type Direction uint8

const (
	Outbound Direction = 0
	Inbound  Direction = 1
)

type WheelchairAccessibility uint8

const (
	NoInfo        WheelchairAccessibility = 0
	Accessible    WheelchairAccessibility = 1
	NotAccessible WheelchairAccessibility = 2
)

type Trip struct {
	Id                   uint32                  `json:"id"`
	RouteId              string                  `json:"routeId"`
	ServiceId            uint32                  `json:"serviceId"`
	TripId               string                  `json:"tripId"`
	TripHeadsign         string                  `json:"tripHeadsign"`
	DirectionId          Direction               `json:"directionId"`
	BlockId              uint32                  `json:"blockId"`
	ShapeId              string                  `json:"shapeId"`
	WheelchairAccessible WheelchairAccessibility `json:"wheelchairAccessible"`
}

type LocationCategory uint8

const (
	StopOrPlatform LocationCategory = 0
	Station        LocationCategory = 1
	EntranceExit   LocationCategory = 2
	GenericNode    LocationCategory = 3
	BoardingArea   LocationCategory = 4
)

type Stop struct {
	Id            uint16           `json:"Id"`
	StopCode      string           `json:"stopCode"`
	StopName      string           `json:"stopName"`
	StopDesc      string           `json:"stopDesc"`
	StopLat       float32          `json:"stopLat"`
	StopLon       float32          `json:"stopLon"`
	StopUrl       string           `json:"stopUrl"`
	LocationType  LocationCategory `json:"locationType"`
	ParentStation uint16           `json:"parentStation"`
}

type PickupStatus uint8

const (
    RegularPickup     PickupStatus = 0
    NoPickup          PickupStatus = 1
    PhoneAgency       PickupStatus = 2
    CoordinateDriver  PickupStatus = 3 
)

type DropOffStatus uint8

const (
    RegularDropOff     DropOffStatus = 0
    NoDropOff          DropOffStatus = 1
    PhoneAgencyDropOff DropOffStatus = 2
    CoordinateDriverDropOff DropOffStatus = 3
)

type StopTime struct {
	Id          	uint32 			`json:"id"`
	TripId      	string 			`json:"tripId"`
	ArrivalTime 	time.Time 		`json:"arrivalTime"`
	DepartureTime 	time.Time 		`json:"departureTime"`
	StopId      	uint16 			`json:"stopId"`
	StopSequence 	uint8 			`json:"stopSequence"`
	PickupType  	PickupStatus 	`json:"pickupType"`
	DropOffType 	DropOffStatus 	`json:"dropOffType"`
}

type Shape struct{
	Id 				uint32 		`json:"id"`
	ShapeId 		string 		`json:"shapeId"`
	ShapePtLat		float32		`json:"shapePtLat"`
	ShapePtLon		float32 	`json:"shapePtLon"`
	ShapePtSequence uint16 		`json:"shapePtSequence"`
	ShapeDistTraveled float32 	`json:"shapeDistTraveled"`
}

type TransportType uint8

const (
    TramStreetcarLightRail TransportType = 0
    SubwayMetro            TransportType = 1
    Rail                   TransportType = 2
    Bus                    TransportType = 3
    Ferry                  TransportType = 4
    CableTram              TransportType = 5
    AerialLift             TransportType = 6
    Funicular              TransportType = 7
    Trolleybus             TransportType = 11
    Monorail               TransportType = 12
)

type Route struct {
	Id 				uint32 			`json:"id"`
	RouteId 		string 			`json:"routeId"`
	RouteShortName 	string 			`json:"routeShortName"`
	RouteLongName 	string 			`json:"routeLongName"`
	RouteDesc 		string 			`json:"routeDesc"`
	RouteType 		TransportType 	`json:"routeType"`
	RouteUrl 		string 			`json:"routeUrl"`
	RouteColor 		string 			`json:"routeColor"`
	RouteTextColor 	string 			`json:"routeTextColor"`
	RouteSortOrder 	uint32 			`json:"routeSortOrder"`
}

type FareRules struct {
	Id 			uint32 	`json:"id"`
	FareId 		uint32 	`json:"fareId"`
	RouteId 	string 	`json:"routeId"`
	OriginId 	uint16 	`json:"originId"`
	DestinationId uint16 `json:"destinationId"`
	ContainsId 	uint16 	`json:"containsId"`
}

type FarePayment uint8

const (
    FareOnBoard         FarePayment = 0
    FareBeforeBoarding  FarePayment = 1
)

type TransferPermitted uint8

const (
    NoTransfers      TransferPermitted = 0 
    OneTransfer      TransferPermitted = 1 
    TwoTransfers     TransferPermitted = 2 
    UnlimitedTransfers TransferPermitted = 255 
)

type FareAttributes struct {
	Id 				uint32 			`json:"id"`
	FareId 			uint32 			`json:"fareId"`
	Price 			float32 		`json:"price"`
	CurrencyType 	string 			`json:"currencyType"`
	PaymentMethod 	FarePayment 	`json:"paymentMethod"`
	Transfers 		TransferPermitted `json:"transfers"`
	TransferDuration uint16 		`json:"transferDuration"`
}

type ServiceAvailability uint8

const (
    ServiceAdded    ServiceAvailability = 1
    ServiceRemoved  ServiceAvailability = 2
)

type CalendarDates struct {
	Id 				uint32 				`json:"id"`
	ServiceId 		uint32 				`json:"serviceId"`
	Date 			time.Time 			`json:"date"`
	ExceptionType 	ServiceAvailability `json:"exceptionType"`
}

type DayServiceAvailability uint8

const (
    ServiceAvailable   DayServiceAvailability = 1
    ServiceNotAvailable DayServiceAvailability = 0
)

type Calendar struct {
	Id 				uint32 								`json:"id"`
	ServiceId 		uint32 								`json:"serviceId"`
	WeekDaysService map[string]DayServiceAvailability 	`json:"weekServices"`
	StartData 		time.Time 							`json:"startDate"`
	EndDate 		time.Time 							`json:"endDate"`
}

type Agency struct {
	Id 				uint8 	`json:"id"`
	AgencyId 		string 	`json:"agencyId"`
	AgencyUrl 		string 	`json:"agencyUrl"`
	AgencyTimezone 	string 	`json:"agencyTimezone"`
	AgencyPhone 	string 	`json:"agencyPhone"`
	AgencyLang 		string 	`json:"agencyLang"`
}
