package models

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type League struct {
	ID        string          `json:"id" gorm:"primary_key"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt gorm.DeletedAt  `json:"deletedAt,omitempty"`
	Name      string          `json:"name"`
	Manual    string          `json:"manual"`
	StartAt   string          `json:"startAt"`
	FinishAt  string          `json:"finishAt"`
	Rules     *Rules          `json:"rules,omitempty"`
	UIDS      []AdminsLeagues `json:"uids"`
}

type AdminsLeagues struct {
	LeagueID  string `json:"leagueId" gorm:"primaryKey;autoIncrement:false"`
	UID       string `json:"uid" gorm:"primaryKey;autoIncrement:false"`
	AdminName string `json:"adminName"`
}

func (l *League) BeforeCreate(tx *gorm.DB) (err error) {
	// uuidをstringに変換し‐を除去
	uid := uuid.New()
	l.ID = strings.Replace(uid.String(), "-", "", -1)
	l.Name = strings.TrimSpace(l.Name)

	if utf8.RuneCountInString(l.Name) < 5 || utf8.RuneCountInString(l.Name) > 30 {
		return errors.New("error")
	}
	if l.Name == "" {
		return errors.New("error")
	}
	if utf8.RuneCountInString(l.Manual) > 100 {
		return errors.New("error")
	}
	l.Manual = strings.TrimSpace(l.Manual)

	return
}
