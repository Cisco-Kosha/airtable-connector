package models

import "time"

type Records struct {
	ID          string    `json:"id"`
	CreatedTime time.Time `json:"createdTime"`
	Fields      struct {
		Field4   string `json:"Field 4"`
		Field3   string `json:"Field 3"`
		Projects string `json:"Projects"`
		EndDate  string `json:"End date"`
	} `json:"fields"`
}
