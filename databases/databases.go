package databases

import(
	"time"
)

// Model Database
type Model struct {
	CreatedAt    		time.Time				`gorm:"type:datetime;"`
	UpdatedAt    		time.Time				`gorm:"type:datetime;"`
	DeleteAt			time.Time				`gorm:"type:datetime;default:NULL"`
}

// Product Schema Database
type Product struct {
	ID					int						`gorm:"primaryKey"`
	Code  				string					`gorm:"type:varchar(255);"`
	Price 				int						`gorm:"type:int(255);"`
	Model

}

// User Schema Database
type User struct {
	ID           		int 					`gorm:"primaryKey"`
	Name         		string					`gorm:"type:varchar(255);"`
	Email        		string					`gorm:"type:varchar(255);"`
	Age          		int						`gorm:"type:int(255);"`
	Birthday     		time.Time				`gorm:"type:datetime;"`
	MemberNumber 		string					`gorm:"type:varchar(255);"`
	Model
}

// Author Schema Database
type Author struct {
	ID					int						`gorm:"primaryKey"`
	Name  				string					`gorm:"type:varchar(255);"`
	Email 				string					`gorm:"type:varchar(255);"`
	Model
}