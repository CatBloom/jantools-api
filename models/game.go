package models

import (
	"time"
)

type Game struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	LeagueId  string     `json:"leagueId"`
	Players   []Player   `json:"players" gorm:"many2many:games_players"`
	Results   []Result   `json:"results" gorm:"foreignKey:GameID;-:migration"`
}
