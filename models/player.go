package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Player struct {
	PlayerId   uint `gorm:"primary_key" json:"playerId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PlayerName string `json:"playerName" gorm:"index:,unique,composite:myname"`
	LeagueId   uint   `json:"leagueId" gorm:"index:,unique,composite:myname"`
}

func (p *Player) BeforeCreate(tx *gorm.DB) (err error) {
	p.PlayerName = "aaa"
	tx.Statement.SetColumn("PlayerName", "aaaaa")
	fmt.Println("通過")
	return nil
}
