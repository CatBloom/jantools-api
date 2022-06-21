package models

import (
	"time"
)

type League struct {
	LeagueId     uint `gorm:"primary_key" json:"leagueId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	LeagueName   string `json:"leagueName"`
	LeagueManual string `json:"leagueManual"`
	UID          string `json:"uid"`
}
