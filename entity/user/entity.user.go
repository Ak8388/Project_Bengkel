package user

import (
	"time"

	"gorm.io/gorm"
)

type Regist struct {
	gorm.Model
	Nama         string `json:"nama" binding:"required" form:"nama"`
	Alamat       string `json:"alamat" binding:"required"`
	TanggalLahir string `json:"tanggal_lahir" binding:"required"`
	Email        string `json:"email" binding:"required,email" gorm:"unique"`
	Username     string `json:"username" binding:"required" gorm:"unique"`
	Password     string `json:"password" binding:"required"`
}

type TransaksiService struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" form:"" json:"id"`
	NamaLengkap    string    `gorm:"type:character varying(200)" binding:"required" form:"" json:"nama_lengkap"`
	Alamat         string    `gorm:"type:character varying(200)" binding:"required" form:"" json:"alamat"`
	NomerHp        string    `gorm:"type:character varying(13)" binding:"required,min=11,max=13,startswith=08" form:"" json:"no_hp"`
	Email          string    `gorm:"type:character varying(100)" binding:"required,email" form:"" json:"email"`
	Tanggal        string    `gorm:"type:character varying(21)" binding:"required,min=17,max=21" form:"" json:"tanggal"`
	TanggalService time.Time `gorm:"type:date" form:"" json:"tanggal_service"`
	Service        string    `gorm:"type:character varying(150)" binding:"required" form:"" json:"service"`
	TotalHarga     uint      `form:"" json:"total_harga"`
	Model          string    `gorm:"type:character varying(150)" binding:"required" form:"" json:"model"`
	TahunModel     string    `gorm:"type:char(4)" binding:"required,min=4,max=4" form:"" json:"tahun_model"`
}

type TransaksiSparepart struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" form:"" json:"id"`
	NamaLengkap string `gorm:"type:character varying(200)" binding:"required" form:"" json:"nama_lengkap"`
	Alamat      string `gorm:"type:character varying(200)" binding:"required" form:"" json:"alamat"`
	NomerHp     string `gorm:"type:character varying(13)" binding:"required,min=11,max=13,startswith=08" form:"" json:"no_hp"`
	Email       string `gorm:"type:character varying(100)" binding:"required,email" form:"" json:"email"`
	SparePart   string `gorm:"type:character varying(150)" binding:"required" form:"" json:"sparepart"`
	Jumlah      uint   `form:"" json:"jumlah" binding:"required"`
	TotalHarga  uint   `form:"" json:"total_harga"`
	Model       string `gorm:"type:character varying(150)" binding:"required" form:"" json:"model"`
	TahunModel  string `gorm:"type:char(4)" binding:"required,min=4,max=4" form:"" json:"tahun_model"`
}
