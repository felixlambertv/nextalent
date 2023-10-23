package model

type Person struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name    string `gorm:"type:varchar(255);NOT NULL"`
	Country string `gorm:"type:varchar(255);NOT NULL"`
}
