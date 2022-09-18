package services

import (
	"fmt"

	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameService interface {
	GetGameList(lid string) ([]models.Game, error)
	GetGame(id string) (models.Game, error)
	CreateGame(c *gin.Context) (models.Game, error)
	UpdateGame(id string, c *gin.Context) (models.Game, error)
	DeleteGame(id string) error
	CheckGameAuth(id string, c *gin.Context) error
}

type gameService struct {
	db *gorm.DB
}

func NewGameService(db *gorm.DB) GameService {
	return &gameService{db}
}

type Game = models.Game

//gamelistの取得
func (s gameService) GetGameList(lid string) ([]models.Game, error) {
	var g []Game

	if err := s.db.Preload("Results", func(db *gorm.DB) *gorm.DB {
		return db.Select("results.*,players.name").Order("Results.rank ASC").Joins("left join players on players.id = results.player_id")
	}).Where("league_id = ?", lid).Order("created_at desc").Find(&g).Error; err != nil {
		return g, err
	}

	return g, nil
}

//gameの取得
func (s gameService) GetGame(id string) (models.Game, error) {
	var g Game

	if err := s.db.Preload("Results", func(db *gorm.DB) *gorm.DB {
		return db.Select("results.*,players.name").Order("Results.rank ASC").Joins("left join players on players.id = results.player_id")
	}).First(&g, "id = ?", id).Error; err != nil {
		return g, err
	}

	return g, nil
}

//gameの作成
func (s gameService) CreateGame(c *gin.Context) (models.Game, error) {
	var g Game

	if err := c.BindJSON(&g); err != nil {
		return g, err
	}

	if err := s.db.Omit("Players.*").Create(&g).Error; err != nil {
		return g, err
	}

	// 作成時playerNameを取得する
	if err := s.db.Preload("Results", func(db *gorm.DB) *gorm.DB {
		return db.Select("*").Order("Results.rank ASC").Joins("left join players on players.id = results.player_id")
	}).Find(&g, "id = ?", g.ID).Error; err != nil {
		return g, err
	}

	return g, nil
}

//gameの更新
func (s gameService) UpdateGame(id string, c *gin.Context) (models.Game, error) {
	var g Game

	//gameの更新はトランザクション処理をする
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := c.BindJSON(&g); err != nil {
			return err
		}
		if err := s.db.Unscoped().Table("results").Where("game_id = ?", id).Delete("").Error; err != nil {
			return err
		}
		if err := s.db.Unscoped().Table("games_players").Where("game_id = ?", id).Delete("").Error; err != nil {
			return err
		}
		if err := s.db.Session(&gorm.Session{FullSaveAssociations: true}).Omit("Players.*").Updates(&g).Error; err != nil {
			return err
		}
		return nil
	})
	//トランザクション処理のエラー判定
	if err != nil {
		return g, err
	}

	return g, nil
}

//gameの削除
func (s gameService) DeleteGame(id string) error {
	var g Game

	//gameの削除はトランザクション処理をする
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.db.Unscoped().Table("results").Where("game_id = ?", id).Delete("").Error; err != nil {
			return nil
		}
		if err := s.db.Unscoped().Table("games_players").Where("game_id = ?", id).Delete("").Error; err != nil {
			return nil
		}

		if err := s.db.Where("id = ?", id).Delete(&g).Error; err != nil {
			return err
		}
		return nil
	})
	//トランザクション処理のエラー判定
	if err != nil {
		return err
	}

	return nil
}

func (s gameService) CheckGameAuth(id string, c *gin.Context) error {
	var g Game

	uid, _ := c.Get("fb_uid")
	sUid := fmt.Sprint(uid)

	if err := s.db.
		Joins("left join admins_leagues as al on games.league_id = al.league_id").
		Where("uid = ? AND games.id = ?", sUid, id).
		First(&g).Error; err != nil {
		return err
	}

	return nil
}
