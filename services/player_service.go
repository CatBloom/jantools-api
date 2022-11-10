package services

import (
	"fmt"

	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlayerService interface {
	GetPlayerList(id string) ([]models.Player, error)
	CreatePlayer(c *gin.Context) (models.Player, error)
	UpdatePlayer(id string, c *gin.Context) (Player, error)
	DeletePlayer(id string) error
	CheckPlayerAuth(id string, c *gin.Context) error
}

type playerService struct {
	db *gorm.DB
}

func NewPlayerService(db *gorm.DB) PlayerService {
	return &playerService{db}
}

type Player = models.Player

//playerlistの取得
func (s playerService) GetPlayerList(id string) ([]models.Player, error) {
	var p []Player

	if err := s.db.Table("players").Where("league_id = ?", id).Order("name ASC, created_at ASC").Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

//playerの作成
func (s playerService) CreatePlayer(c *gin.Context) (models.Player, error) {
	var p Player

	if err := c.BindJSON(&p); err != nil {
		return p, err
	}
	if err := s.db.Table("players").Create(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

//playerの更新
func (s playerService) UpdatePlayer(id string, c *gin.Context) (Player, error) {
	var p Player

	if err := s.db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	if err := c.BindJSON(&p); err != nil {
		return p, err
	}

	s.db.Save(&p)

	return p, nil
}

//playerの削除
func (s playerService) DeletePlayer(id string) error {
	var p Player

	if err := s.db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}

	return nil
}

func (s playerService) CheckPlayerAuth(id string, c *gin.Context) error {
	var p Player

	uid, _ := c.Get("fb_uid")
	sUid := fmt.Sprint(uid)

	if err := s.db.
		Joins("left join admins_leagues as al on players.league_id = al.league_id").
		Where("uid = ? AND players.id = ?", sUid, id).
		First(&p).Error; err != nil {
		return err
	}

	return nil
}
