package model

type Tes struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"notNull"`
	Description string `json:"description"`
}
