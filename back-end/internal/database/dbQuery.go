package database

import (
	"busProject/models"
	"gorm.io/gorm"
)

func GetRoutesWithDays(db *gorm.DB) ([]models.RouteWithDays, error) {
	var results []models.RouteWithDays

	rows, err := db.Raw(`
		SELECT 
			r.route_id, r.route_short_name, r.route_long_name, r.route_type, 
			r.route_url, r.route_color, r.route_text_color, r.route_sort_order,
			MAX(c.monday) AS monday, MAX(c.tuesday) AS tuesday, MAX(c.wednesday) AS wednesday, 
			MAX(c.thursday) AS thursday, MAX(c.friday) AS friday, MAX(c.saturday) AS saturday, 
			MAX(c.sunday) AS sunday
		FROM routes r
		JOIN trips t ON r.route_id = t.route_id
		JOIN calendars c ON t.service_id = c.service_id
		GROUP BY r.route_id
	`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var route models.RouteWithDays
		var monday, tuesday, wednesday, thursday, friday, saturday, sunday int

		if err := rows.Scan(
			&route.RouteID, &route.RouteShortName, &route.RouteLongName, &route.RouteType,
			&route.RouteURL, &route.RouteColor, &route.RouteTextColor, &route.RouteSortOrder,
			&monday, &tuesday, &wednesday, &thursday, &friday, &saturday, &sunday,
		); err != nil {
			return nil, err
		}

		// Convert binary days into weekday names
		weekdays := []string{}
		if monday == 1 {
			weekdays = append(weekdays, "P")
		}
		if tuesday == 1 {
			weekdays = append(weekdays, "A")
		}
		if wednesday == 1 {
			weekdays = append(weekdays, "T")
		}
		if thursday == 1 {
			weekdays = append(weekdays, "K")
		}
		if friday == 1 {
			weekdays = append(weekdays, "P")
		}
		if saturday == 1 {
			weekdays = append(weekdays, "Š")
		}
		if sunday == 1 {
			weekdays = append(weekdays, "S")
		}

		route.WeekDays = weekdays
		results = append(results, route)
	}

	return results, nil
}