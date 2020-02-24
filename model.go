package main

import "time"

func insertRow(testing *Testing) {
	if err := db.Create(&testing).Error; err != nil {
		panic(err)
	}
}

//Testing storage
type Testing struct {
	FirstName   string
	LastName    string
	Phone       string
	LastUpdated time.Time
}
