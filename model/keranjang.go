package model

import (
	"time"
)

type Keranjang struct {
	ID           uint `gorm:"primarykey" json:"primarykey"`
	KasirID      uint
	Kasir        User `gorm:"foreignkey:KasirID"`
	ProductID    uint
	Product      Product `gorm:"foreignkey:ProductID"`
	JumlahBarang uint
	TotalHarga   uint
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// di kasir bisa buat nota. satu nota bisa
