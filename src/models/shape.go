package models

type Shape struct{
	Id 				uint32 		`json:"id"`
	ShapeId 		string 		`json:"shapeId"`
	ShapePtLat		float32		`json:"shapePtLat"`
	ShapePtLon		float32 	`json:"shapePtLon"`
	ShapePtSequence uint16 		`json:"shapePtSequence"`
	ShapeDistTraveled float32 	`json:"shapeDistTraveled"`
}