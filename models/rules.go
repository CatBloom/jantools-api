package models

import (
	"time"

	"gorm.io/gorm"
)

type Rules struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
	LeagueId    string     `json:"league_id" gorm:"foreignkey"`
	PlayerCount int        `json:"playerCount"`
	GameType    string     `json:"gameType"`
	Tanyao      bool       `json:"tanyao"`
	Back        bool       `json:"back"`
	Dora        int        `json:"dora"`
	StartPoint  int        `json:"startPoint"`
	ReturnPoint int        `json:"returnPoint"`
	Uma1        int        `json:"uma1"`
	Uma2        int        `json:"uma2"`
	Uma3        int        `json:"uma3"`
	Uma4        *int       `json:"uma4"`
}

func (rules *Rules) BeforeCreate(tx *gorm.DB) (err error) {
	if rules.PlayerCount == 3 {
		rules.Uma4 = nil
	}
	return
}
