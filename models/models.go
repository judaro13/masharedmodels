package models

import "time"

// QueryCoordinatesMessage message struct
type QueryCoordinatesMessage struct {
	Reference   string     `json:"reference"`
	Coordinates [][]string `json:"coordinates"`
}

// StoreDataMessage message struct
type StoreDataMessage struct {
	Reference string          `json:"reference"`
	Result    UKAPIPOSTResult `json:"result"`
}

// UKAPIBulkQuery struct to generate body body
type UKAPIBulkQuery struct {
	Geolocations []UKAPICoordinate `json:"geolocations"`
}

// UKAPICoordinate API coordintate struct
type UKAPICoordinate struct {
	Postcode  string  `json:"postcode"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Radius    int     `json:"radius"`
	Limit     int     `json:"limit"`
}

// UKAPIPOSTResult post query result
type UKAPIPOSTResult struct {
	Status int            `json:"status"`
	Result []UKAPIResults `json:"result"`
}

// UKAPIResults API single solution struct
type UKAPIResults struct {
	Result []UKAPICoordinate `json:"result"`
}

// GeoCoordinate db struct
type GeoCoordinate struct {
	ID        uint `gorm:"primaryKey;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Lat       float64 `gorm:"UNIQUE_INDEX:compositeindex;index;not null" json:"lat"`
	Lon       float64 `gorm:"UNIQUE_INDEX:compositeindex;index;not null" json:"lon"`
	Postcode  string  `gorm:"not null" json:"postcode"`
}

// CSVUpload db struct to handle progress
type CSVUpload struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Records   int `json:"-"`
	Bulks     int
	Counts    int
	Status    string
	Reference string
}
