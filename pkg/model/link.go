package model

import "time"

type Link struct {
	Id        int
	Link      string
	CropLink  string
	IpAddress string
	CreatedAt time.Time
}
