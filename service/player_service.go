package service

import (
	"github.com/CatBloom/MahjongMasterApi/db"
	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

// Service procides Player's behavior
type PlayerService struct {
	// db *gorm.DB
}

// func NewPlayerService(db *gorm.DB) *PlayerService {
// 	return &PlayerService{db}
// }

// Player is alias of entity.Player struct
type Player models.Player

// CreateModel is create Player model
func (s PlayerService) CreateModel(c *gin.Context) (Player, error) {
	db := db.GetDB()
	var u Player

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s PlayerService) GetByLeagueID(id string) ([]Player, error) {
	db := db.GetDB()
	var u []Player

	if err := db.Where("league_id = ?", id).Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// UpdateByID is update a Player
func (s PlayerService) UpdateByID(id string, c *gin.Context) (Player, error) {
	db := db.GetDB()
	var u Player

	if err := db.Where("player_id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

// DeleteByID is delete a Player
func (s PlayerService) DeleteByID(id string) error {
	db := db.GetDB()
	var u Player

	if err := db.Where("player_id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
