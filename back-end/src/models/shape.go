package models

type Shape struct{
	ShapeId 			string 		`json:"shapeId"`
	ShapePtLat			float64		`json:"shapePtLat"`
	ShapePtLon			float64 	`json:"shapePtLon"`
	ShapePtSequence 	int 		`json:"shapePtSequence"`
	ShapeDistTraveled 	int 		`json:"shapeDistTraveled"`
}