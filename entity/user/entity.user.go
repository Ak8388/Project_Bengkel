package user

import "time"

type TransaksiService struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaLengkap    string    `gorm:"type:character varying(200)" binding:"required" json:"nama_lengkap"`
	Alamat         string    `gorm:"type:character varying(200)" binding:"required" json:"alamat"`
	NomerHp        string    `gorm:"type:character varying(13)" binding:"required,min=11,max=13,startswith=08" json:"no_hp"`
	Email          string    `gorm:"type:character varying(100)" binding:"required,email" json:"email"`
	Tanggal        string    `gorm:"type:character varying(21)" binding:"required,min=17,max=21" json:"tanggal"`
	TanggalService time.Time `gorm:"type:date" json:"tanggal_service"`
	Service        string    `gorm:"type:character varying(150)" binding:"required" json:"service"`
	TotalHarga     uint      `json:"total_harga"`
	Model          string    `gorm:"type:character varying(150)" binding:"required" json:"model"`
	TahunModel     string    `gorm:"type:char(4)" binding:"required,min=4,max=4" json:"tahun_model"`
}

type TransaksiSparepart struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaLengkap string `gorm:"type:character varying(200)" binding:"required" json:"nama_lengkap"`
	Alamat      string `gorm:"type:character varying(200)" binding:"required" json:"alamat"`
	NomerHp     string `gorm:"type:character varying(13)" binding:"required,min=11,max=13,startswith=08" json:"no_hp"`
	Email       string `gorm:"type:character varying(100)" binding:"required,email" json:"email"`
	SparePart   string `gorm:"type:character varying(150)" binding:"required" json:"sparepart"`
	Jumlah      uint   `json:"jumlah" binding:"required"`
	TotalHarga  uint   `json:"total_harga"`
	Model       string `gorm:"type:character varying(150)" binding:"required" json:"model"`
	TahunModel  string `gorm:"type:char(4)" binding:"required,min=4,max=4" json:"tahun_model"`
}
