package domain

type Points []*Point

type Point struct {
	Name string  `json:"name" bson:"name"`
	Lat  float32 `json:"lattitude" bson:"lattitude,truncate"`
	Long float32 `json:"longitude" bson:"longitude,truncate"`
}
