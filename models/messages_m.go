package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type Messages struct {
// 	ID        string    `gorm:"primaryKey;type:varchar(36);uniqueIndex" json:"uuid"`
// 	Author    string    `gorm:"type:varchar(64)" json:"author"`
// 	Message   string    `gorm:"type:varchar(1024)" json:"message"`
// 	Likes     int64     `gorm:"" json:"likes"`
// 	UpdateAt  time.Time `gorm:"autoUpdateTime" json:"-"`
// 	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
// }

type Member struct {
	// ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	gorm.Model
	Name      string `gorm:"type:varchar(64)" json:"name"`
	TelNumber string `gorm:"varchar(10);unique" json:"tel"`
	// Record    []RentRecord `gorm:""`
}

type Field struct {
	// ID    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	gorm.Model
	Name  string `gorm:"type:varchar(64);unique" json:"name"`
	Price int    `json:"price"`
}

type RentRecord struct {
	// ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	gorm.Model
	MemberID  uuid.UUID `json:"member_id"`
	Member    Member    `gorm:"foreignKey:MemberID" json:"member"`
	Start     time.Time `json:"start_time"`
	End       time.Time `json:"end_time"`
	Field     Field     `gorm:"foreignKey:FieldName;references:Name"`
	FieldName string    `gorm:"type:varchar(64)" json:"field_name"`
}
