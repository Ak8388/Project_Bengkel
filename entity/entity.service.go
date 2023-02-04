package entity

type Service struct {
	ID             uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	TypeService    string           `gorm:"type:character varying(200)" binding:"required" json:"type_service"`
	ServiceUmum    []ServiceUmum    `gorm:"foreignKey:id_service"`
	ServiceBerkala []ServiceBerkala `gorm:"foreignKey:id_service"`
}

type ServiceUmum struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Service   string `gorm:"type:character varying(200)" json:"service"`
	Harga     uint   `binding:"required" json:"harga"`
	IdService uint   `json:"id_service" binding:"required"`
}

type ServiceBerkala struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Service   string `gorm:"type:character varying(200)" json:"service"`
	Harga     uint   `binding:"required" json:"harga"`
	IdService uint   `json:"id_service" binding:"required"`
}
