package services

import (
	"fmt"

	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LeagueService interface {
	SearchLeague(value string) ([]League, error)
	GetLeagueList(c *gin.Context) ([]League, error)
	GetLeague(id string) (League, error)
	CreateLeague(c *gin.Context) (League, error)
	UpdateLeague(id string, c *gin.Context) (League, error)
	DeleteLeague(id string) error
}
type leagueService struct {
	db *gorm.DB
}

func NewLeagueService(db *gorm.DB) LeagueService {
	return &leagueService{db}
}

type League = models.League

//league検索用
func (s leagueService) SearchLeague(value string) ([]League, error) {
	var l []League
	if err := s.db.Where("name Like ?", value+"%").Order("created_at desc").Limit(5).Find(&l).Error; err != nil {
		return nil, err
	}

	return l, nil
}

//leaguelistの取得
func (s leagueService) GetLeagueList(c *gin.Context) ([]League, error) {
	var l []League
	uid, _ := c.Get("fb_uid")
	sUid := fmt.Sprint(uid)
	if err := s.db.Preload("Rules").Joins("left join admins_leagues as al on leagues.id = al.league_id").Where("uid = ?", sUid).Order("created_at desc").Find(&l).Error; err != nil {
		return nil, err
	}

	return l, nil
}

//leagueの取得
func (s leagueService) GetLeague(id string) (League, error) {
	var l League

	if err := s.db.Preload("UIDS").Preload("Rules").Table("leagues").Find(&l, "leagues.id = ?", id).Error; err != nil {
		return l, err
	}

	return l, nil
}

//leagueの作成
func (s leagueService) CreateLeague(c *gin.Context) (League, error) {
	var l League

	uid, _ := c.Get("fb_uid")
	name, _ := c.Get("fb_name")
	sUid := fmt.Sprint(uid)
	sName := fmt.Sprint(name)

	//leagueの登録はトランザクション処理をする
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := c.BindJSON(&l); err != nil {
			return err
		}

		if err := tx.Create(&l).Error; err != nil {
			return err
		}

		a2l := &models.AdminsLeagues{
			LeagueID:  l.ID,
			UID:       sUid,
			AdminName: sName,
		}

		if err := tx.Create(&a2l).Error; err != nil {
			return err
		}

		return nil
	})
	//トランザクション処理のエラー判定
	if err != nil {
		return l, err
	}

	return l, nil
}

//leagueの更新
func (s leagueService) UpdateLeague(id string, c *gin.Context) (League, error) {
	var l League

	if err := s.db.Where("id = ?", id).First(&l).Error; err != nil {
		return l, err
	}
	if err := c.BindJSON(&l); err != nil {
		return l, err
	}

	s.db.Save(&l)

	return l, nil
}

//leagueの削除
func (s leagueService) DeleteLeague(id string) error {
	var l League

	//leagueの削除は論理削除
	if err := s.db.Debug().Where("id = ?", id).Delete(&l).Error; err != nil {
		return err
	}
	return nil
}
