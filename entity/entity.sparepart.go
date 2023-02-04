package entity

type Kategori struct {
	ID        uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	KaName    string      `gorm:"type:character varying(100)" binding:"required" json:"name"`
	SparePart []SparePart `gorm:"foreignKey:id_kat"`
}

type SparePart struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama    string `gorm:"type:character varying(200)" binding:"required" json:"nama"`
	Harga   uint   `json:"harga" binding:"required"`
	Stok    uint   `json:"stok" binding:"required"`
	IdKat   uint   `json:"id_kat" binding:"required"`
	IdModel uint   `json:"id_model" binding:"required"`
}
