package entity

type MerkKendaraan struct {
	ID             uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	Merk           string           `gorm:"type:character varying(200)" binding:"required" json:"merk"`
	ModelKendaraan []ModelKendaraan `gorm:"foreignKey:id_merk"`
}

type ModelKendaraan struct {
	ID        uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	Model     string      `gorm:"type:character varying(150)" binding:"required" json:"model"`
	Tahun     string      `gorm:"type:char(4)" binding:"required" json:"tahun"`
	IdMerk    uint        `json:"id_merk" binding:"required"`
	SparePart []SparePart `gorm:"foreignKey:id_model"`
}
