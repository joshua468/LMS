package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	Amount    int64  `json:"amount"`
	Status    string `json:"status"`
}
