package models

import (
	"time"
)

type Player struct {
	PlayerId   uint `gorm:"primary_key" json:"playerId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PlayerName string `json:"playerName"`
	LeagueId   uint   `json:"leagueId"`
}
