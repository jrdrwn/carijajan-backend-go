package main

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type Pedagang struct {
	ID            int            `json:"id"`
	Nama          string         `json:"nama"`
	JenisDagangan string         `json:"jenis_dagangan"`
	FotoUrl       string         `json:"foto_url"`
	JamBuka       datatypes.Time `json:"jam_buka"`
	JamTutup      datatypes.Time `json:"jam_tutup"`
	Status        string         `json:"status"`
	Lokasi        datatypes.JSON `json:"lokasi"`
}

type PedagangWithJarak struct {
	Pedagang
	JarakKM float64 `json:"jarak_km"`
}

type PedagangDetail struct {
	Pedagang
	Deskripsi   string         `json:"deskripsi"`
	Tags        datatypes.JSON `json:"tags"`
	NoHp        string         `json:"no_hp"`
	HariMangkal pq.StringArray `json:"hari_mangkal" gorm:"type:text[]"`
	CreatedAt   time.Time      `json:"created_at"`
}

type JenisDagangan struct {
	Value string `json:"jenis_dagangan" gorm:"jenis_dagangan"`
}
