package model2

import "time"

type Transaction struct {
	ID        uint
	HewanId   uint
	PembeliId uint
	Tanggal   time.Time
	Total     float64
	Status    string
}
