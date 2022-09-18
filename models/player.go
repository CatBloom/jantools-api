package models

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"gorm.io/gorm"
)

type Player struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name"`
	LeagueId  string     `json:"leagueId"`
	Games     []Game     `json:"games" gorm:"many2many:games_players"`
}

func (p *Player) BeforeCreate(tx *gorm.DB) (err error) {
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" {
		return errors.New("error")
	}
	if utf8.RuneCountInString(p.Name) > 10 {
		return errors.New("error")
	}
	return
}

func (p *Player) BeforeUpdate(tx *gorm.DB) (err error) {
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" {
		return errors.New("error")
	}
	if utf8.RuneCountInString(p.Name) > 10 {
		return errors.New("error")
	}
	return
}
