package models

type Prize struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text"`
	Stock       int    `json:"stock" gorm:"default:0"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime"`
}

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Email     string `json:"email" gorm:"type:varchar(255);unique;not null"`
	Phone     string `json:"phone" gorm:"type:varchar(20);unique"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime"`
}

type LotteryEntry struct {
	ID        uint  `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint  `json:"user_id" gorm:"not null"`
	PrizeID   uint  `json:"prize_id" gorm:"not null"`
	Timestamp int64 `json:"timestamp" gorm:"autoCreateTime"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`

	// 外键关联(搜了一下hhh,原来要这样干吗？)
	User  User  `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Prize Prize `json:"prize" gorm:"foreignKey:PrizeID;references:ID"`
}
