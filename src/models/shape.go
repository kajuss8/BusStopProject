package models

type Shape struct{
	Id 				int 		`json:"id"`
	ShapeId 		string 		`json:"shapeId"`
	ShapePtLat		float32		`json:"shapePtLat"`
	ShapePtLon		float32 	`json:"shapePtLon"`
	ShapePtSequence int 		`json:"shapePtSequence"`
	ShapeDistTraveled int 		`json:"shapeDistTraveled"`
}