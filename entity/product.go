package entity

import "time"

type Product struct {
	ID        int
	Name      string
	Price     int
	Desc      string
	Qty       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
