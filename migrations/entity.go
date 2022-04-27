package migrations

import (
	"time"
)

type Buku struct {
	ID          int    `gorm:"primary_key"`
	IdKategori  int    `gorm:"not null"`
	Judul       string `gorm:"type:varchar(100);not null"`
	Pengarang   string `gorm:"type:varchar(100);not null"`
	Penerbit    string `gorm:"type:varchar(100);not null"`
	TahunTerbit int    `gorm:"not null"`
	JumlahBuku  int    `gorm:"not null"`
	Harga       int    `gorm:"not null"`
	Gambar      string `gorm:"type:varchar(255);not null"`
	Deskripsi   string `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Cart struct {
	ID        int `gorm:"primary_key"`
	IdBuku    int `gorm:"not null"`
	Jumlah    int `gorm:"not null"`
	SubTotal  int `gorm:"not null"`
	IdDiskon  int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DetailTransaction struct {
	IdTransaction int `gorm:"not null"`
	IdBuku        int `gorm:"not null"`
	Jumlah        int `gorm:"not null"`
	SubTotal      int `gorm:"not null"`
	IdDiskon      int `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Diskon struct {
	ID           int `gorm:"primary_key"`
	IdBuku       int `gorm:"not null"`
	Diskon       int `gorm:"not null"`
	MinPembelian int `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Kategori struct {
	ID           int    `gorm:"primary_key"`
	NamaKategori string `gorm:"type:varchar(100);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Transaction struct {
	ID             int       `gorm:"primary_key"`
	TglTransaction time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	JamTransaction time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Total          int       `gorm:"not null"`
	Bayar          int       `gorm:"not null"`
	Item           int       `gorm:"not null"`
	Kembali        int       `gorm:"not null"`
	IdUser         int       `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type User struct {
	ID        int    `gorm:"primary_key"`
	Nama      string `gorm:"type:varchar(255);not null"`
	Username  string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Token     string
	ExpToken  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
